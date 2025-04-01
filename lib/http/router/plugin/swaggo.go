package plugin

import (
	librouter "tns-energo/lib/http/router"

	"github.com/swaggo/http-swagger/v2"
)

type Swaggo struct {
	serviceName string
}

func NewSwaggo(serviceName string) *Swaggo {
	return &Swaggo{
		serviceName: serviceName,
	}
}

func (s *Swaggo) BasePath() string {
	return "/debug/swagger"
}

func (s *Swaggo) Register(router *librouter.Router) {
	basePath := "/debug/swagger/doc.json"
	if len(s.serviceName) > 0 {
		basePath = "/" + s.serviceName + basePath
	}

	cfg := httpSwagger.URL(basePath)
	handler := httpSwagger.Handler(cfg)

	router.PathPrefix("").Handler(librouter.WrapStdLibFunc(handler))
}
