package document

import (
	"io"
	libctx "tns-energo/lib/ctx"
)

type Repository interface {
	Create(ctx libctx.Context, fileName string, payload io.Reader, payloadLength int64) (string, error)
}
