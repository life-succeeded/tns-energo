package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"tns-energo/lib/apierr"
	librouter "tns-energo/lib/http/router"
)

func Recover(c librouter.Context, h librouter.Handler) librouter.Handler {
	log := c.Log()

	return func(c librouter.Context) error {
		var (
			rq = c.Request()
			id = fmt.Sprintf("%s %s", rq.Method, rq.URL.Path)
		)

		defer func() {
			if panicErr := recover(); panicErr != nil {
				log.Errorf("Паника при выполнении %s: %+v\n%s", id, panicErr, debug.Stack())
				_ = c.WriteJson(http.StatusInternalServerError,
					librouter.NewErrorResponse(apierr.CodePanic, "Something went wrong, we have a panic!"))
			}
		}()

		return h(c)
	}
}
