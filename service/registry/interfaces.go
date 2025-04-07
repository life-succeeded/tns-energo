package registry

import (
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
)

type RegistryStorage interface {
	AddMany(ctx libctx.Context, items []Item) error
	GetByAccountNumber(ctx libctx.Context, accountNumber string) (Item, error)
	GetByAccountNumberRegular(ctx libctx.Context, log liblog.Logger, accountNumber string) ([]Item, error)
}
