package main

import (
	"context"
	"fmt"
	"io/fs"
	"time"
	"tns-energo/api"
	"tns-energo/config"
	dbuser "tns-energo/database/user"
	"tns-energo/lib/ctx"
	"tns-energo/lib/db"
	libserver "tns-energo/lib/http/server"
	liblog "tns-energo/lib/log"
	"tns-energo/service/user"

	"github.com/jmoiron/sqlx"
)

const _databaseTimeout = 15 * time.Second

type App struct {
	/* main */
	mainCtx  context.Context
	log      liblog.Logger
	settings config.Settings

	/* storage */
	postgres *sqlx.DB

	/* http */
	server libserver.Server

	/* services */
	userService user.Service
}

func NewApp(mainCtx ctx.Context, log liblog.Logger, settings config.Settings) *App {
	return &App{
		mainCtx:  mainCtx,
		log:      log,
		settings: settings,
	}
}

func (a *App) InitDatabases(fs fs.FS, migrationPath string) (err error) {
	postgresCtx, cancelPostgresCtx := context.WithTimeout(a.mainCtx, _databaseTimeout)
	defer cancelPostgresCtx()

	if a.postgres, err = db.NewPgx(postgresCtx, a.settings.Databases.Postgres); err != nil {
		return fmt.Errorf("could not connect to postgres: %w", err)
	}

	if err = db.Migrate(fs, a.log, a.postgres, migrationPath); err != nil {
		return fmt.Errorf("could not migrate postgres: %w", err)
	}

	return nil
}

func (a *App) InitServices() (err error) {
	userRepository := dbuser.NewRepository(a.postgres)

	a.userService = user.NewService(userRepository)

	return nil
}

func (a *App) InitServer() {
	sb := api.NewServerBuilder(a.mainCtx, a.log, a.settings)
	sb.AddDebug()
	sb.AddUsers(a.userService)
	a.server = sb.Build()
}

func (a *App) Start() {
	a.server.Start()
}

func (a *App) Stop(ctx context.Context) {
	a.server.Stop()

	if err := a.postgres.Close(); err != nil {
		a.log.Errorf("could not close postgres connection: %v", err)
	}
}
