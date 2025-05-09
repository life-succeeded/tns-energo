package object

import (
	"context"
	"fmt"
	"io"
	libctx "tns-energo/lib/ctx"
	"tns-energo/lib/db"

	"github.com/minio/minio-go/v7"
)

type Minio struct {
	mc           *minio.Client
	bucket, host string
}

func NewStorage(ctx context.Context, mc *minio.Client, bucket, host string) (*Minio, error) {
	err := db.PrepareBucket(ctx, mc, bucket)
	if err != nil {
		return &Minio{}, fmt.Errorf("could not prepare bucket: %w", err)
	}

	return &Minio{
		mc:     mc,
		bucket: bucket,
		host:   host,
	}, nil
}

func (s *Minio) Add(ctx libctx.Context, fileName string, payload io.Reader, payloadLength int64) (string, error) {
	_, err := s.mc.PutObject(ctx, s.bucket, fileName, payload, payloadLength, minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("could not put object %s: %w", fileName, err)
	}

	return fmt.Sprintf("%s/%s/%s", s.host, s.bucket, fileName), nil
}
