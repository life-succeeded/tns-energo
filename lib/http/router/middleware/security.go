package middleware

import (
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
// Проверяет указана ли в запросе валидная авторизация с признаком админа. В случае провала проверки - вызывает failHandler, если он указан
func EnableCors(_ librouter.Context, h librouter.Handler) librouter.Handler {
	return func(c librouter.Context) error {
		c.WriteHeader("Access-Control-Allow-Origin", "*")
		c.WriteHeader("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.WriteHeader("Access-Control-Allow-Headers", "*")

		return h(c)
	}
}
