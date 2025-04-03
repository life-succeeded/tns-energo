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
	"tns-energo/lib/db/minio"
	libserver "tns-energo/lib/http/server"
	liblog "tns-energo/lib/log"
	"tns-energo/service/inspection"
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
	minio    minio.Client

	/* http */
	server libserver.Server

	/* services */
	userService       user.Service
	inspectionService inspection.Service
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

	minioCtx, cancelMinioCtx := context.WithTimeout(a.mainCtx, _databaseTimeout)
	defer cancelMinioCtx()

	if a.minio, err = minio.NewClient(
		minioCtx,
		a.settings.Databases.Minio.Endpoint,
		a.settings.Databases.Minio.User,
		a.settings.Databases.Minio.Password,
		a.settings.Databases.Minio.UseSSL,
		[]string{a.settings.Databases.Minio.ImagesBucket, a.settings.Databases.Minio.DocumentsBucket},
		a.settings.Databases.Minio.Host,
	); err != nil {
		return fmt.Errorf("could not connect to minio: %w", err)
	}

	return nil
}

func (a *App) InitServices() (err error) {
	userRepository := dbuser.NewRepository(a.postgres)

	a.userService = user.NewService(userRepository, a.settings)
	a.inspectionService = inspection.NewService(a.minio, a.settings)

	return nil
}

func (a *App) InitServer() {
	sb := api.NewServerBuilder(a.mainCtx, a.log, a.settings)
	sb.AddDebug()
	sb.AddUsers(a.userService)
	sb.AddInspections(a.inspectionService)
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
