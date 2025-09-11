package analytics

import (
	"io"
	"time"
	"tns-energo/service/inspection"

	"github.com/sunshineOfficial/golib/goctx"
	"github.com/sunshineOfficial/golib/golog"
)

type ReportStorage interface {
	AddOne(ctx goctx.Context, report Report) (string, error)
	GetAll(ctx goctx.Context, log golog.Logger) ([]Report, error)
}

type InspectionStorage interface {
	GetByInspectionDate(ctx goctx.Context, log golog.Logger, inspectionDate time.Time) ([]inspection.Inspection, error)
}

type DocumentStorage interface {
	Add(ctx goctx.Context, fileName string, payload io.Reader, payloadLength int64) (string, error)
}
