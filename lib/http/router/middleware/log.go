package middleware

import (
	"time"
	liberr "tns-energo/lib/err"
	librouter "tns-energo/lib/http/router"
)

func LogError(_ librouter.Context, h librouter.Handler) librouter.Handler {
	return func(c librouter.Context) error {
		var (
			log = c.Log()
			err = h(c)
		)

		switch {
		case err == nil:
			return nil

		case liberr.IsCriticalError(err):
			log.Error(err.Error())

		default:
			log.Debug(err.Error())
		}

		return err
	}
}

func VerboseLog(_ librouter.Context, h librouter.Handler) librouter.Handler {
	return func(c librouter.Context) error {
		start := time.Now()

		defer func() {
			var (
				rq    = c.Request()
				rs    = c.Response()
				log   = c.Log()
				total = time.Since(start)
			)

			if rs.IsCommitted() {
				log.Debugf("Запрос %s %s выполнен за %s", rq.Method, rq.URL.Path, total)
			} else {
				log.Debugf("Запрос %s %s (статус %d) выполнен за %s", rq.Method, rq.URL.Path, rs.Status(), total)
			}
		}()

		return h(c)
	}
}
