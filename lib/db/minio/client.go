package minio

import (
	"context"
	"fmt"
	libctx "tns-energo/lib/ctx"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client interface {
	CreateOne(ctx libctx.Context, bucket string, file File) error
}

type Impl struct {
	mc *minio.Client
}

func NewClient(ctx context.Context, endpoint, user, password string, useSSL bool, buckets []string) (*Impl, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("could not create minio client: %w", err)
	}

	for _, bucket := range buckets {
		exists, err := client.BucketExists(ctx, bucket)
		if err != nil {
			return nil, fmt.Errorf("could not check if bucket exists: %w", err)
		}

		if exists {
			continue
		}

		err = client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("could not create bucket %s: %w", bucket, err)
		}
	}

	return &Impl{
		mc: client,
	}, nil
}

func (c *Impl) CreateOne(ctx libctx.Context, bucket string, file File) error {
	_, err := c.mc.PutObject(ctx, bucket, file.Name, file.Data, int64(file.Data.Len()), minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("could not put file %s: %w", file.Name, err)
	}

	return nil
}
