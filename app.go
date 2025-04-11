package main

import (
	"context"
	"fmt"
	"time"
	"tns-energo/api"
	"tns-energo/config"
	dbbrigade "tns-energo/database/brigade"
	dbinspection "tns-energo/database/inspection"
	"tns-energo/database/object"
	dbregistry "tns-energo/database/registry"
	dbreport "tns-energo/database/report"
	dbtask "tns-energo/database/task"
	"tns-energo/lib/ctx"
	"tns-energo/lib/db"
	libserver "tns-energo/lib/http/server"
	liblog "tns-energo/lib/log"
	"tns-energo/service/analytics"
	"tns-energo/service/brigade"
	"tns-energo/service/cron"
	"tns-energo/service/image"
	"tns-energo/service/inspection"
	"tns-energo/service/registry"
	"tns-energo/service/task"

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
	mongo *mongo.Client
	minio *minio.Client

	/* http */
	server libserver.Server

	/* services */
	inspectionService *inspection.Service
	registryService   *registry.Service
	imageService      *image.Service
	analyticsService  *analytics.Service
	cronService       *cron.Service
	taskService       *task.Service
	brigadeService    *brigade.Service
}

func NewApp(mainCtx ctx.Context, log liblog.Logger, settings config.Settings) *App {
	return &App{
		mainCtx:  mainCtx,
		log:      log,
		settings: settings,
	}
}

func (a *App) InitDatabases() (err error) {
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
	inspectionStorage := dbinspection.NewStorage(a.mongo, a.settings.Inspections.Database, a.settings.Inspections.Collection)
	registryStorage := dbregistry.NewStorage(a.mongo, a.settings.Registry.Database, a.settings.Registry.Collection)
	reportStorage := dbreport.NewStorage(a.mongo, a.settings.Reports.Database, a.settings.Reports.Collection)
	taskStorage := dbtask.NewStorage(a.mongo, a.settings.Tasks.Database, a.settings.Tasks.Collection)
	brigadeStorage := dbbrigade.NewStorage(a.mongo, a.settings.Brigades.Database, a.settings.Brigades.Collection)

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

	a.inspectionService = inspection.NewService(a.settings, inspectionStorage, documentStorage, registryStorage, taskStorage, brigadeStorage)
	a.registryService = registry.NewService(registryStorage)
	a.imageService = image.NewService(imageStorage)
	a.analyticsService = analytics.NewService(reportStorage)
	a.cronService = cron.NewService(a.settings, a.analyticsService)
	a.taskService = task.NewService(taskStorage)
	a.brigadeService = brigade.NewService(brigadeStorage)

	return nil
}

func (a *App) InitServer() {
	sb := api.NewServerBuilder(a.mainCtx, a.log, a.settings)
	sb.AddDebug()
	sb.AddInspections(a.inspectionService)
	sb.AddRegistry(a.registryService)
	sb.AddImages(a.imageService)
	sb.AddAnalytics(a.analyticsService)
	sb.AddTasks(a.taskService)
	sb.AddBrigades(a.brigadeService)
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

	mongoCtx, cancelMongoCtx := context.WithTimeout(ctx, _databaseTimeout)
	defer cancelMongoCtx()

	if err := a.mongo.Disconnect(mongoCtx); err != nil {
		a.log.Errorf("could not close mongodb connection: %v", err)
	}
}
