package analytics

import (
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
)

type ReportStorage interface {
	AddOne(ctx libctx.Context, report Report) (string, error)
	GetAll(ctx libctx.Context, log liblog.Logger) ([]Report, error)
}
