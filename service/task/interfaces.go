package task

import (
	"github.com/sunshineOfficial/golib/goctx"
	"github.com/sunshineOfficial/golib/golog"
)

type Storage interface {
	AddOne(ctx goctx.Context, task Task) (string, error)
	GetByBrigadeId(ctx goctx.Context, log golog.Logger, brigadeId string) ([]Task, error)
	UpdateStatus(ctx goctx.Context, id string, status Status) error
	GetById(ctx goctx.Context, id string) (Task, error)
}
