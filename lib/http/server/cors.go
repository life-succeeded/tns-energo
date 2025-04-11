package server

import (
	"github.com/gorilla/handlers"
	"net/http"
	librouter "tns-energo/lib/http/router"
)

func WithCors(router *librouter.Router) http.Handler {
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"*"}),
	)

	return corsMiddleware(router)
}
