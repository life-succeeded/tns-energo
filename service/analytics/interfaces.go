package analytics

import libctx "tns-energo/lib/ctx"

type ReportStorage interface {
	AddOne(ctx libctx.Context, report Report) error
}
