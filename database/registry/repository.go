package registry

import libctx "tns-energo/lib/ctx"

type Repository interface {
	AddOne(ctx libctx.Context, item Item) error
	AddMany(ctx libctx.Context, items []Item) error
	GetByAccountNumber(ctx libctx.Context, accountNumber string) (Item, error)
}
