package brigade

import (
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
)

type Storage interface {
	AddOne(ctx libctx.Context, b Brigade) (string, error)
	GetById(ctx libctx.Context, id string) (Brigade, error)
	GetAll(ctx libctx.Context, log liblog.Logger) ([]Brigade, error)
	Update(ctx libctx.Context, id string, b Brigade) error
}
