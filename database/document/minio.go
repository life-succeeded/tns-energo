package document

import (
	"context"
	"fmt"
	"io"
	libctx "tns-energo/lib/ctx"
	"tns-energo/lib/db"

	"github.com/minio/minio-go/v7"
)

type Minio struct {
	mc     *minio.Client
	bucket string
	host   string
}

func NewRepository(ctx context.Context, mc *minio.Client, bucket, host string) (Minio, error) {
	err := db.PrepareBucket(ctx, mc, bucket)
	if err != nil {
		return Minio{}, fmt.Errorf("could not prepare bucket: %w", err)
	}

	return Minio{
		mc:     mc,
		bucket: bucket,
		host:   host,
	}, nil
}

func (r Minio) Create(ctx libctx.Context, fileName string, payload io.Reader, payloadLength int64) (string, error) {
	_, err := r.mc.PutObject(ctx, r.bucket, fileName, payload, payloadLength, minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("could not put object %s: %w", fileName, err)
	}

	return fmt.Sprintf("%s/%s/%s", r.host, r.bucket, fileName), nil
}
