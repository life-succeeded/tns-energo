package http

import (
	"net/http"
	"strings"
	"tns-energo/lib/authorize"
	libctx "tns-energo/lib/ctx"
	"tns-energo/lib/locale"
	liblog "tns-energo/lib/log"

	"github.com/google/uuid"
)

func GetContext(log liblog.Logger, r *http.Request) libctx.Context {
	token := authorize.GetToken(r)
	ctx := libctx.Context{
		Context:   r.Context(),
		AuthToken: token,
		Locale:    locale.Get(r),
		Origin:    r.Header.Get(OriginHeader),
	}

	requestId := r.Header.Get(RequestIdHeader)
	if len(requestId) == 0 {
		ctx.RequestId = uuid.New()
	} else {
		id, err := uuid.Parse(requestId)
		if err == nil {
			ctx.RequestId = id
		}
	}

	if len(token) > 0 && strings.Index(token, "Bearer ") == 0 {
		auth, err := authorize.Parse(token)
		if err == nil {
			ctx.Authorize = auth
		} else {
			log.Debugf("Некорректный токен авторизации: %v", err)
		}
	}

	return ctx
}

func RequestWithContext(r *http.Request, c libctx.Context) *http.Request {
	if len(c.AuthToken) > 0 {
		r.Header.Set(AuthorizationHeader, c.AuthToken)
	}
	if len(c.Locale) > 0 {
		locale.Set(r, c.Locale)
	}
	if len(c.Origin) > 0 {
		r.Header.Set(OriginHeader, c.Origin)
	}
	if stringId := c.RequestId.String(); len(stringId) > 0 {
		r.Header.Set(RequestIdHeader, stringId)
	}

	return r
}
