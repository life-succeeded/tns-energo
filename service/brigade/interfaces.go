package brigade

import libctx "tns-energo/lib/ctx"

type Storage interface {
	AddOne(ctx libctx.Context, b Brigade) (string, error)
	GetById(ctx libctx.Context, id string) (Brigade, error)
}
