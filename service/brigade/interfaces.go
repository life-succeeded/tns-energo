package brigade

import (
	"github.com/sunshineOfficial/golib/goctx"
	"github.com/sunshineOfficial/golib/golog"
)

type Storage interface {
	AddOne(ctx goctx.Context, b Brigade) (string, error)
	GetById(ctx goctx.Context, id string) (Brigade, error)
	GetAll(ctx goctx.Context, log golog.Logger) ([]Brigade, error)
	Update(ctx goctx.Context, id string, b Brigade) error
}
