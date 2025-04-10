package task

import (
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
)

type Storage interface {
	AddOne(ctx libctx.Context, task Task) (string, error)
	GetByInspectorId(ctx libctx.Context, log liblog.Logger, inspectorId int) ([]Task, error)
	UpdateStatus(ctx libctx.Context, id string, status Status) error
}
