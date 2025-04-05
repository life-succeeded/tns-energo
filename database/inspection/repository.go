package inspection

import libctx "tns-energo/lib/ctx"

type Repository interface {
	CreateOne(ctx libctx.Context, inspection Inspection) error
	CreateMany(ctx libctx.Context, inspections []Inspection) error
}
