package inspection

import (
	"io"
	libctx "tns-energo/lib/ctx"
)

type DocumentStorage interface {
	Add(ctx libctx.Context, fileName string, payload io.Reader, payloadLength int) (string, error)
}
