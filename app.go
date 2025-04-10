package main

import (
	"context"
	"fmt"
	"io/fs"
	"time"
	"tns-energo/api"
	"tns-energo/config"
	dbinspection "tns-energo/database/inspection"
	"tns-energo/database/object"
	dbregistry "tns-energo/database/registry"
	dbreport "tns-energo/database/report"
	dbtask "tns-energo/database/task"
	dbuser "tns-energo/database/user"
	"tns-energo/lib/ctx"
	"tns-energo/lib/db"
	libserver "tns-energo/lib/http/server"
	liblog "tns-energo/lib/log"
	"tns-energo/service/analytics"
	"tns-energo/service/cron"
	"tns-energo/service/image"
	"tns-energo/service/inspection"
	"tns-energo/service/registry"
	"tns-energo/service/task"
	"tns-energo/service/user"

	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.mongodb.org/mongo-driver/mongo"
)

const _databaseTimeout = 15 * time.Second

type App struct {
	/* main */
	mainCtx  context.Context
	log      liblog.Logger
	settings config.Settings

	/* storage */
	postgres *sqlx.DB
	mongo    *mongo.Client
	minio    *minio.Client

	/* http */
	server libserver.Server

	/* services */
	userService       *user.Service
	inspectionService *inspection.Service
	registryService   *registry.Service
	imageService      *image.Service
	analyticsService  *analytics.Service
	cronService       *cron.Service
	taskService       *task.Service
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

	mongoCtx, cancelMongoCtx := context.WithTimeout(a.mainCtx, _databaseTimeout)
	defer cancelMongoCtx()

	if a.mongo, err = db.NewMongo(mongoCtx, a.settings.Databases.Mongo); err != nil {
		return fmt.Errorf("could not connect to mongodb: %w", err)
	}

	if a.minio, err = minio.New(a.settings.Databases.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(a.settings.Databases.Minio.User, a.settings.Databases.Minio.Password, ""),
		Secure: a.settings.Databases.Minio.UseSSL,
	}); err != nil {
		return fmt.Errorf("could not connect to minio: %w", err)
	}

	return nil
}

func (a *App) InitServices() (err error) {
	userStorage := dbuser.NewStorage(a.postgres)
	inspectionStorage := dbinspection.NewStorage(a.mongo, a.settings.Inspections.Database, a.settings.Inspections.Collection)
	registryStorage := dbregistry.NewStorage(a.mongo, a.settings.Registry.Database, a.settings.Registry.Collection)
	reportStorage := dbreport.NewStorage(a.mongo, a.settings.Reports.Database, a.settings.Reports.Collection)
	taskStorage := dbtask.NewStorage(a.mongo, a.settings.Tasks.Database, a.settings.Tasks.Collection)

	documentCtx, cancelDocumentCtx := context.WithTimeout(a.mainCtx, _databaseTimeout)
	defer cancelDocumentCtx()

	documentStorage, err := object.NewStorage(documentCtx, a.minio, a.settings.Databases.Minio.DocumentsBucket, a.settings.Databases.Minio.Host)
	if err != nil {
		return fmt.Errorf("could not create document storage: %w", err)
	}

	imageCtx, cancelImageCtx := context.WithTimeout(a.mainCtx, _databaseTimeout)
	defer cancelImageCtx()

	imageStorage, err := object.NewStorage(imageCtx, a.minio, a.settings.Databases.Minio.ImagesBucket, a.settings.Databases.Minio.Host)
	if err != nil {
		return fmt.Errorf("could not create image storage: %w", err)
	}

	a.userService = user.NewService(a.settings, userStorage)
	a.inspectionService = inspection.NewService(a.settings, inspectionStorage, documentStorage, userStorage, registryStorage, taskStorage)
	a.registryService = registry.NewService(registryStorage)
	a.imageService = image.NewService(imageStorage)
	a.analyticsService = analytics.NewService(reportStorage)
	a.cronService = cron.NewService(a.settings, a.analyticsService)
	a.taskService = task.NewService(taskStorage)

	return nil
}

func (a *App) InitServer() {
	sb := api.NewServerBuilder(a.mainCtx, a.log, a.settings)
	sb.AddDebug()
	sb.AddUsers(a.userService)
	sb.AddInspections(a.inspectionService)
	sb.AddRegistry(a.registryService)
	sb.AddImages(a.imageService)
	sb.AddAnalytics(a.analyticsService)
	sb.AddTasks(a.taskService)
	a.server = sb.Build()
}

func (a *App) Start() error {
	a.server.Start()

	if err := a.cronService.LaunchJobs(a.mainCtx, a.log); err != nil {
		return fmt.Errorf("could not launch jobs: %w", err)
	}

	return nil
}

func (a *App) Stop(ctx context.Context) {
	if err := a.cronService.Shutdown(); err != nil {
		a.log.Errorf("could not shutdown cron scheduler: %v", err)
	}

	a.server.Stop()

	if err := a.postgres.Close(); err != nil {
		a.log.Errorf("could not close postgres connection: %v", err)
	}

	mongoCtx, cancelMongoCtx := context.WithTimeout(ctx, _databaseTimeout)
	defer cancelMongoCtx()

	if err := a.mongo.Disconnect(mongoCtx); err != nil {
		a.log.Errorf("could not close mongodb connection: %v", err)
	}
}
