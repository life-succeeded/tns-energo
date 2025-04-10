package api

import (
	"context"
	"fmt"
	"tns-energo/api/handlers"
	"tns-energo/config"
	librouter "tns-energo/lib/http/router"
	"tns-energo/lib/http/router/middleware"
	"tns-energo/lib/http/router/plugin"
	"tns-energo/lib/http/router/status"
	libserver "tns-energo/lib/http/server"
	liblog "tns-energo/lib/log"
	"tns-energo/service/analytics"
	"tns-energo/service/image"
	"tns-energo/service/inspection"
	"tns-energo/service/registry"
	"tns-energo/service/user"
)

type ServerBuilder struct {
	server       libserver.Server
	router       *librouter.Router
	isAuthorized librouter.Middleware
}

func NewServerBuilder(ctx context.Context, log liblog.Logger, settings config.Settings) *ServerBuilder {
	return &ServerBuilder{
		server:       libserver.NewHTTPServer(ctx, log, fmt.Sprintf(":%d", settings.Port)),
		router:       librouter.NewRouter(log).Use(middleware.Recover, middleware.LogError),
		isAuthorized: middleware.IsAnyAuthorized(status.ForbiddenHandler),
	}
}

func (s *ServerBuilder) AddDebug() {
	s.router.Install(plugin.NewPProf())
}

func (s *ServerBuilder) AddUsers(userService *user.Service) {
	subRouter := s.router.SubRouter("/users")
	subRouter.HandlePost("/register", handlers.RegisterHandler(userService))
	subRouter.HandlePost("/login", handlers.LoginHandler(userService))
	subRouter.HandlePut("/refresh-token/{refresh_token}", handlers.RefreshTokenHandler(userService))
	subRouter.HandleGet("/{user_id}", handlers.GetUserByIdHandler(userService))
}

func (s *ServerBuilder) AddInspections(inspectionService *inspection.Service) {
	subRouter := s.router.SubRouter("/inspections")
	subRouter.HandlePost("", handlers.InspectHandler(inspectionService)).Use(s.isAuthorized)
	subRouter.HandleGet("/by-inspector-id/{inspector_id}", handlers.GetInspectionsByInspectorId(inspectionService))
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

func (s *ServerBuilder) Build() libserver.Server {
	s.server.UseHandler(s.router)

	return s.server
}
