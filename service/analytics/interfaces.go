package analytics

import (
	"io"
	"time"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	"tns-energo/service/inspection"
)

type ReportStorage interface {
	AddOne(ctx libctx.Context, report Report) (string, error)
	GetAll(ctx libctx.Context, log liblog.Logger) ([]Report, error)
}

type InspectionStorage interface {
	GetByInspectionDate(ctx libctx.Context, log liblog.Logger, inspectionDate time.Time) ([]inspection.Inspection, error)
}

type DocumentStorage interface {
	Add(ctx libctx.Context, fileName string, payload io.Reader, payloadLength int64) (string, error)
}
