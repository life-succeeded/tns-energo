package inspection

import (
	"io"
	"time"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	"tns-energo/service/brigade"
	"tns-energo/service/registry"
	"tns-energo/service/task"
)

type Storage interface {
	AddOne(ctx libctx.Context, inspection Inspection) error
	GetByBrigadeId(ctx libctx.Context, log liblog.Logger, brigadeId string) ([]Inspection, error)
	GetByInspectionDate(ctx libctx.Context, log liblog.Logger, inspectionDate time.Time) ([]Inspection, error)
}

type DocumentStorage interface {
	Add(ctx libctx.Context, fileName string, payload io.Reader, payloadLength int64) (string, error)
}

type RegistryStorage interface {
	GetByAccountNumber(ctx libctx.Context, accountNumber string) (registry.Item, error)
	AddOne(ctx libctx.Context, item registry.Item) error
	UpdateOne(ctx libctx.Context, item registry.Item) error
}

type TaskStorage interface {
	UpdateStatus(ctx libctx.Context, id string, status task.Status) error
}

type BrigadeStorage interface {
	GetById(ctx libctx.Context, id string) (brigade.Brigade, error)
}
