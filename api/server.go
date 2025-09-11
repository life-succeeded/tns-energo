package api

import (
	"context"
	"fmt"
	"tns-energo/api/handlers"
	"tns-energo/config"
	"tns-energo/service/analytics"
	"tns-energo/service/brigade"
	"tns-energo/service/image"
	"tns-energo/service/inspection"
	"tns-energo/service/registry"
	"tns-energo/service/task"

	"github.com/sunshineOfficial/golib/gohttp/gorouter"
	"github.com/sunshineOfficial/golib/gohttp/gorouter/middleware"
	"github.com/sunshineOfficial/golib/gohttp/gorouter/plugin"
	"github.com/sunshineOfficial/golib/gohttp/goserver"
	"github.com/sunshineOfficial/golib/golog"
)

type ServerBuilder struct {
	server goserver.Server
	router *gorouter.Router
}

func NewServerBuilder(ctx context.Context, log golog.Logger, settings config.Settings) *ServerBuilder {
	return &ServerBuilder{
		server: goserver.NewHTTPServer(ctx, log, fmt.Sprintf(":%d", settings.Port)),
		router: gorouter.NewRouter(log).Use(
			middleware.Metrics(),
			middleware.Recover,
			middleware.LogError,
		),
	}
}

func (s *ServerBuilder) AddDebug() {
	s.router.Install(plugin.NewPProf(), plugin.NewMetrics())
}

func (s *ServerBuilder) AddInspections(inspectionService *inspection.Service) {
	subRouter := s.router.SubRouter("/inspections")
	subRouter.HandlePost("", handlers.InspectHandler(inspectionService))
	subRouter.HandleGet("/by-brigade-id/{brigade_id}", handlers.GetInspectionsByBrigadeId(inspectionService))
}

func (s *ServerBuilder) AddRegistry(registryService *registry.Service) {
	subRouter := s.router.SubRouter("/registry")
	subRouter.HandlePost("/parse", handlers.ParseRegistryHandler(registryService))
	subRouter.HandleGet("/item/by-account-number/{account_number}", handlers.GetItemByAccountNumberHandler(registryService))
	subRouter.HandleGet("/items/by-account-number/{account_number}/regular", handlers.GetItemsByAccountNumberRegularHandler(registryService))
}

func (s *ServerBuilder) AddImages(imageService *image.Service) {
	subRouter := s.router.SubRouter("/images")
	subRouter.HandlePost("", handlers.UploadImageHandler(imageService))
}

func (s *ServerBuilder) AddAnalytics(analyticsService *analytics.Service) {
	subRouter := s.router.SubRouter("/analytics")
	subRouter.HandlePost("/reports/daily/{date}", handlers.GenerateDailyReportHandler(analyticsService))
	subRouter.HandleGet("/reports", handlers.GetAllReportsHandler(analyticsService))
}

func (s *ServerBuilder) AddTasks(taskService *task.Service) {
	subRouter := s.router.SubRouter("/tasks")
	subRouter.HandlePost("", handlers.AddTaskHandler(taskService))
	subRouter.HandleGet("/by-brigade-id/{brigade_id}", handlers.GetTasksByBrigadeId(taskService))
	subRouter.HandlePatch("/{task_id}/status", handlers.UpdateTaskStatusHandler(taskService))
	subRouter.HandleGet("/{task_id}", handlers.GetTaskById(taskService))
}

func (s *ServerBuilder) AddBrigades(brigadeService *brigade.Service) {
	subRouter := s.router.SubRouter("/brigades")
	subRouter.HandlePost("", handlers.CreateBrigadeHandler(brigadeService))
	subRouter.HandleGet("/{brigade_id}", handlers.GetBrigadeByIdHandler(brigadeService))
	subRouter.HandleGet("", handlers.GetAllBrigadesHandler(brigadeService))
	subRouter.HandlePut("/{brigade_id}", handlers.UpdateBrigadeHandler(brigadeService))
}

func (s *ServerBuilder) Build() goserver.Server {
	s.server.UseHandler(s.router)

	return s.server
}
