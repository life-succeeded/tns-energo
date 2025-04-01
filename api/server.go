package api

import (
	"context"
	"fmt"
	"tns-energo/config"
	librouter "tns-energo/lib/http/router"
	"tns-energo/lib/http/router/middleware"
	"tns-energo/lib/http/router/plugin"
	"tns-energo/lib/http/router/status"
	libserver "tns-energo/lib/http/server"
	liblog "tns-energo/lib/log"
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

func (s *ServerBuilder) Build() libserver.Server {
	s.server.UseHandler(s.router)

	return s.server
}
