package inspection

import (
	"io"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	"tns-energo/service/user"
)

type Storage interface {
	AddOne(ctx libctx.Context, inspection Inspection) error
	GetByInspectorId(ctx libctx.Context, log liblog.Logger, inspectorId int) ([]Inspection, error)
}

type DocumentStorage interface {
	Add(ctx libctx.Context, fileName string, payload io.Reader, payloadLength int) (string, error)
}

type UserStorage interface {
	GetLightById(ctx libctx.Context, userId int) (user.UserLight, error)
}
