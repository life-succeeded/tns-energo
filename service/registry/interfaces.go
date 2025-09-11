package registry

import (
	"github.com/sunshineOfficial/golib/goctx"
	"github.com/sunshineOfficial/golib/golog"
)

type Storage interface {
	AddMany(ctx goctx.Context, items []Item) error
	GetByAccountNumber(ctx goctx.Context, accountNumber string) (Item, error)
	GetByAccountNumberRegular(ctx goctx.Context, log golog.Logger, accountNumber string) ([]Item, error)
}
