package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	liblog "tns-energo/lib/log"
)

const (
	listenTimeout   = 3 * time.Second
	shutdownTimeout = 5 * time.Second
)

type Server interface {
	Start()
	Stop()
	UseHandler(http.Handler)
}

type HTTPServer struct {
	log     liblog.Logger
	ctx     context.Context
	server  *http.Server
	running *atomic.Bool
}

func NewHTTPServer(mainCtx context.Context, log liblog.Logger, addr string) *HTTPServer {
	return &HTTPServer{
		log: log.WithTags("server"),
		ctx: mainCtx,
		server: &http.Server{
			Addr: addr,
			BaseContext: func(_ net.Listener) context.Context {
				return mainCtx
			},
		},
		running: &atomic.Bool{},
	}
}

func (h *HTTPServer) Start() {
	if h.running.Load() {
		return
	}

	h.running.Store(true)
	go h.listen()
}

func (h *HTTPServer) listen() {
	h.log.Debugf("Сервер запускается")

	for h.running.Load() {
		err := h.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			h.log.Debugf("Не удалось запустить сервер: %v. Повторная попытка через %s", err, listenTimeout)
			continue
		}
	}

	h.log.Debugf("Сервер остановлен")
}

func (h *HTTPServer) Stop() {
	if !h.running.Load() {
		return
	}

	h.running.Store(false)

	shutdownCtx, cancel := context.WithTimeout(h.ctx, shutdownTimeout)
	defer cancel()

	if err := h.server.Shutdown(shutdownCtx); err != nil {
		h.log.Debugf("Не удалось остановить сервер за %s: %v", shutdownTimeout, err)
	}
}

func (h *HTTPServer) UseHandler(handler http.Handler) {
	h.server.Handler = handler
}
