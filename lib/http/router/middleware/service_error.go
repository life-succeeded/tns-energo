package middleware

import (
	"net/http"
	liberr "tns-energo/lib/err"
	librouter "tns-energo/lib/http/router"
)

func WithServiceErrors(status int) librouter.Middleware {
	if status == 0 {
		status = http.StatusBadRequest
	}

	return func(_ librouter.Context, h librouter.Handler) librouter.Handler {
		return func(c librouter.Context) error {
			err := h(c)

			serviceErr := liberr.AsServiceError(err)
			if serviceErr == nil {
				return err
			}

			code, message := serviceErr.Info()

			writeErr := c.WriteJson(status, librouter.ErrorResponse{
				Error: librouter.ErrorInfo{
					Code:    code,
					Message: message,
				},
			})
			if writeErr != nil {
				c.Log().Debugf("не удалось записать ошибку в ответ: %v", writeErr)
			}

			return serviceErr.Internal()
		}
	}
}
