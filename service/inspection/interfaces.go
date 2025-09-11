package inspection

import (
	"io"
	"time"
	"tns-energo/service/brigade"
	"tns-energo/service/registry"
	"tns-energo/service/task"

	"github.com/sunshineOfficial/golib/goctx"
	"github.com/sunshineOfficial/golib/golog"
)

type Storage interface {
	AddOne(ctx goctx.Context, inspection Inspection) error
	GetByBrigadeId(ctx goctx.Context, log golog.Logger, brigadeId string) ([]Inspection, error)
	GetByInspectionDate(ctx goctx.Context, log golog.Logger, inspectionDate time.Time) ([]Inspection, error)
}

type DocumentStorage interface {
	Add(ctx goctx.Context, fileName string, payload io.Reader, payloadLength int64) (string, error)
}

type RegistryStorage interface {
	GetByAccountNumber(ctx goctx.Context, accountNumber string) (registry.Item, error)
	AddOne(ctx goctx.Context, item registry.Item) error
	UpdateOne(ctx goctx.Context, item registry.Item) error
}

type TaskStorage interface {
	UpdateStatus(ctx goctx.Context, id string, status task.Status) error
}

type BrigadeStorage interface {
	GetById(ctx goctx.Context, id string) (brigade.Brigade, error)
}
