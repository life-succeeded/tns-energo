package task

import (
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
)

type Storage interface {
	AddOne(ctx libctx.Context, task Task) (string, error)
	GetByBrigadeId(ctx libctx.Context, log liblog.Logger, brigadeId string) ([]Task, error)
	UpdateStatus(ctx libctx.Context, id string, status Status) error
	GetById(ctx libctx.Context, id string) (Task, error)
}
