package minio

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	libctx "tns-energo/lib/ctx"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client interface {
	CreateOne(ctx libctx.Context, bucket string, file File) (string, error)
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

		policy := Policy{
			Version: "2012-10-17",
			Statement: []Statement{
				{
					Effect: "Allow",
					Principal: Principal{
						AWS: []string{"*"},
					},
					Action:   []string{"s3:GetObject"},
					Resource: []string{fmt.Sprintf("arn:aws:s3:::%s/*", bucket)},
				},
			},
		}
		jsonPolicy, err := json.Marshal(policy)
		if err != nil {
			return nil, fmt.Errorf("could not marshal policy: %w", err)
		}

		err = client.SetBucketPolicy(ctx, bucket, string(jsonPolicy))
		if err != nil {
			return nil, fmt.Errorf("could not set bucket policy: %w", err)
		}
	}

	return &Impl{
		mc: client,
	}, nil
}

func (c *Impl) CreateOne(ctx libctx.Context, bucket string, file File) (string, error) {
	_, err := c.mc.PutObject(ctx, bucket, file.Name, file.Data, int64(file.Data.Len()), minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("could not put file %s: %w", file.Name, err)
	}

	url, err := c.mc.PresignedGetObject(ctx, bucket, file.Name, 24*time.Hour, nil)
	if err != nil {
		return "", fmt.Errorf("could not get url for file %s: %w", file.Name, err)
	}

	return url.String(), nil
}
