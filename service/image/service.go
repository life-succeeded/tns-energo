package image

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"time"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	libtime "tns-energo/lib/time"
	"tns-energo/service/file"
)

var imageNumbersCache = make(map[string]int, 32)

type Service struct {
	images Storage
}

func NewService(images Storage) *Service {
	return &Service{
		images: images,
	}
}

func (s *Service) Upload(ctx libctx.Context, log liblog.Logger, request UploadRequest) (file.File, error) {
	payload, err := base64.StdEncoding.DecodeString(request.Payload)
	if err != nil {
		return file.File{}, fmt.Errorf("failed to decode payload: %w", err)
	}

	imageNumbersCache[request.DeviceNumber] = imageNumbersCache[request.DeviceNumber] + 1
	name := fmt.Sprintf("%s_%s_%d.png", request.Address, time.Now().In(libtime.MoscowLocation()).Format("02.01.2006_15.04"), imageNumbersCache[request.DeviceNumber])
	url, err := s.images.Add(ctx, name, bytes.NewReader(payload), len(payload))
	if err != nil {
		return file.File{}, fmt.Errorf("failed to upload image: %w", err)
	}

	return file.File{
		Name: name,
		URL:  url,
	}, nil
}
