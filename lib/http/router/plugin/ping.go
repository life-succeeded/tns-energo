package plugin

import (
	"net/http"
	librouter "tns-energo/lib/http/router"
)

type Ping struct {
}

func NewPing() *Ping {
	return &Ping{}
}

func (p *Ping) BasePath() string {
	return "/debug/ping"
}

func (p *Ping) Register(router *librouter.Router) {
	router.HandleGet("", func(c librouter.Context) error {
		return c.WriteText(http.StatusOK, "pong")
	})
}
