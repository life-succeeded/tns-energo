package router

import (
	"encoding/json"
	"io"
	"net/http"
	libctx "tns-energo/lib/ctx"
	libhttp "tns-energo/lib/http"
	"tns-energo/lib/http/router/reflect"
	liblog "tns-energo/lib/log"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/trace"
)

type Context struct {
	response ResponseWriter
	request  *http.Request

	ctx libctx.Context
	log liblog.Logger

	tracer trace.Tracer
}

func NewContext(log liblog.Logger, rs ResponseWriter, rq *http.Request) Context {
	ctx, _ := libctx.GetContext(rq)
	if ctx.IsAuthorized() {
		log = log.WithUserInfo(ctx.Authorize.UserId, ctx.Authorize.Email)
	}

	return Context{
		response: rs,
		request:  rq,
		ctx:      ctx,
		log:      log,
	}
}

func (c Context) Log() liblog.Logger {
	return c.log
}

func (c Context) Ctx() libctx.Context {
	return c.ctx
}

// CheckAuthorization
// Проверяет авторизацию по указанным правилам:
// - если авторизация не указана или невалиден токен - записывает ответ 401 и возвращает false
// - если указан флаг requireAdmin и в авторизации отсутствует признак админа - записывает ответ 403 и возвращает false
// в остальных случаях - не записывает ничего и возвращает true
func (c Context) CheckAuthorization(requireAdmin bool) bool {
	if !c.ctx.IsAuthorized() {
		c.Write(http.StatusUnauthorized)
		return false
	}
	if requireAdmin && !c.ctx.Authorize.IsAdmin {
		c.Write(http.StatusForbidden)
		return false
	}

	return true
}

func (c Context) Vars(a any) error {
	header := c.request.Header
	query := c.request.URL.Query()

	if err := reflect.SetValuesToItem(getPathVars(c.request), "path", a); err != nil {
		return err
	}
	if err := reflect.SetValuesToItem(header, "header", a); err != nil {
		return err
	}
	if err := reflect.SetValuesToItem(query, "query", a); err != nil {
		return err
	}

	return nil
}

func (c Context) ReadJson(a any) error {
	return json.NewDecoder(c.request.Body).Decode(a)
}

func (c Context) ReadText() (string, error) {
	b, err := io.ReadAll(c.request.Body)
	if err == nil {
		return "", err
	}

	return string(b), nil
}

func (c Context) ReadBytes() ([]byte, error) {
	return io.ReadAll(c.request.Body)
}

func (c Context) Reader() io.ReadCloser {
	return c.request.Body
}

func (c Context) Write(code int) {
	c.response.WriteHeader(code)
}

func (c Context) WriteJson(code int, a any) error {
	return libhttp.WriteResponseJson(c.response, code, a)
}

func (c Context) WriteXml(code int, a any) error {
	return libhttp.WriteResponseXml(c.response, code, a)
}

func (c Context) WriteText(code int, s string) error {
	c.response.WriteHeader(code)
	_, err := io.WriteString(c.response, s)
	return err
}

func (c Context) WriteBinary(code int, b []byte) error {
	c.response.WriteHeader(code)
	_, err := c.response.Write(b)
	return err
}

func (c Context) WriteHeader(name, value string) {
	c.response.Header().Set(name, value)
}

func (c Context) Writer() io.WriteCloser {
	return c.response
}

func (c Context) Request() *http.Request {
	return c.request
}

func (c Context) Response() ResponseWriter {
	return c.response
}

// Tracer возвращает trace.Tracer, ассоциированный с swrouter
// ВАЖНО! Не стоит использовать этот трейсер для трассировки бизнес-логики, т. к. он ассоциирован именно с роутингом.
// Для бизнес-логики используйте трейсеры, созданные в app.go
func (c Context) Tracer() trace.Tracer {
	return c.tracer
}

func (c Context) Close() error {
	if c.request.Body != nil {
		_ = c.request.Body.Close()
	}

	return c.response.Close()
}

func getPathVars(r *http.Request) map[string][]string {
	vars := mux.Vars(r)
	if len(vars) == 0 {
		return nil
	}

	result := make(map[string][]string, len(vars))
	for key, value := range vars {
		result[key] = []string{value}
	}

	return result
}
