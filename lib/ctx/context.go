package ctx

import (
	"context"
	"time"
	"tns-energo/lib/authorize"

	"github.com/google/uuid"
)

type Context struct {
	context.Context

	AuthToken string
	Locale    string
	Origin    string
	Authorize authorize.Authorize
	RequestId uuid.UUID
}

func (c Context) IsAuthorized() bool {
	return len(c.AuthToken) > 0 && c.Authorize.UserId > 0 && len(c.Authorize.Email) > 0
}

func (c Context) IsDone() bool {
	select {
	case <-c.Done():
		return true
	default:
		return false
	}
}

func (c Context) WithCancel() (Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(c.Context)
	return c.cloneWith(ctx), cancel
}

func (c Context) WithTimeout(timeout time.Duration) (Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Context, timeout)
	return c.cloneWith(ctx), cancel
}

func (c Context) WithDeadline(timeout time.Time) (Context, context.CancelFunc) {
	ctx, cancel := context.WithDeadline(c.Context, timeout)
	return c.cloneWith(ctx), cancel
}

// CloneTo возвращает Context, созданный на основе пользовательских данных из контекста, у которого вызывается этот метод.
// Полученный контекст будет независим от таймаутов / дедлайнов и отмен контекста, у которого метод был вызван, но унаследует все его пользовательские данные
func (c Context) CloneTo(ctx context.Context) Context {
	return c.cloneWith(ctx)
}

func (c Context) cloneWith(ctx context.Context) Context {
	return Context{
		Context:   ctx,
		AuthToken: c.AuthToken,
		Authorize: c.Authorize,
		Locale:    c.Locale,
		Origin:    c.Origin,
		RequestId: c.RequestId,
	}
}

func Wrap(ctx context.Context) Context {
	return Context{
		Context: ctx,
	}
}

func Background() Context {
	return Wrap(context.Background())
}
