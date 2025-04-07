package image

import (
	"bytes"
	"fmt"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"

	"github.com/google/uuid"
)

type Service struct {
	images Storage
}

func NewService(images Storage) *Service {
	return &Service{
		images: images,
	}
}

func (s *Service) Upload(ctx libctx.Context, log liblog.Logger, payload []byte) (Image, error) {
	name := fmt.Sprintf("%s.png", uuid.New()) // TODO: указывать расширение в зависимости от типа картинки
	url, err := s.images.Add(ctx, name, bytes.NewReader(payload), len(payload))
	if err != nil {
		return Image{}, fmt.Errorf("failed to upload image: %w", err)
	}

	return Image{
		Name: name,
		URL:  url,
	}, nil
}
