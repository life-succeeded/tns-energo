package image

import (
	"io"

	"github.com/sunshineOfficial/golib/goctx"
)

type Storage interface {
	Add(ctx goctx.Context, fileName string, payload io.Reader, payloadLength int64) (string, error)
}
