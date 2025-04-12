package image

import (
	"io"
	libctx "tns-energo/lib/ctx"
)

type Storage interface {
	Add(ctx libctx.Context, fileName string, payload io.Reader, payloadLength int64) (string, error)
}
