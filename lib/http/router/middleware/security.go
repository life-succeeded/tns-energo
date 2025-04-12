package middleware

import (
	"net/http"
	librouter "tns-energo/lib/http/router"
	"tns-energo/lib/http/router/status"
)

// IsAnyAuthorized
// Проверяет указана ли в запросе любая валидная авторизация. В случае провала проверки - вызывает failHandler, если он указан
func IsAnyAuthorized(failHandler librouter.Handler) librouter.Middleware {
	if failHandler == nil {
		failHandler = status.ForbiddenHandler
	}

	return func(c librouter.Context, h librouter.Handler) librouter.Handler {
		if !c.CheckAuthorization(false) {
			return failHandler
		}

		return h
	}
}

// IsAdmin
// Проверяет указана ли в запросе валидная авторизация с признаком админа. В случае провала проверки - вызывает failHandler, если он указан
func IsAdmin(failHandler librouter.Handler) librouter.Middleware {
	if failHandler == nil {
		failHandler = status.ForbiddenHandler
	}

	return func(c librouter.Context, h librouter.Handler) librouter.Handler {
		if !c.CheckAuthorization(true) {
			return failHandler
		}

		return h
	}
}

// EnableCors
// Включает Cors
func EnableCors(up librouter.Context, h librouter.Handler) librouter.Handler {
	return func(c librouter.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, HEAD, PATCH, CONNECT, TRACE")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		if c.Request().Method == http.MethodOptions {
			c.Response().WriteHeader(http.StatusNoContent)
		}

		return h(c)
	}
}
