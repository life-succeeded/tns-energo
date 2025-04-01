package ctx

import (
	"net/http"
	"tns-energo/lib/authorize"
	"tns-energo/lib/locale"

	"github.com/google/uuid"
)

const (
	AuthorizationHeader = "Authorization"
	OriginHeader        = "Origin"
	RequestIdHeader     = "X-Request-Id"
)

func GetContext(r *http.Request) (Context, error) {
	token := authorize.GetToken(r)
	ctx := Context{
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

	if len(token) > 0 {
		auth, err := authorize.Parse(token)
		if err != nil {
			return ctx, err
		}

		ctx.Authorize = auth
	}

	return ctx, nil
}

func (c Context) WriteHeaders(r *http.Request) {
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
}
