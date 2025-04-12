package image

import (
	"fmt"
	"strings"
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

func (s *Service) Upload(ctx libctx.Context, _ liblog.Logger, request UploadRequest) (file.File, error) {
	payload, err := request.Payload.Open()
	if err != nil {
		return file.File{}, fmt.Errorf("failed to open payload: %w", err)
	}

	split := strings.Split(request.Payload.Filename, ".")
	extension := split[len(split)-1]

	imageNumbersCache[request.DeviceNumber] = imageNumbersCache[request.DeviceNumber] + 1
	name := fmt.Sprintf("%s_%s_%d.%s", request.Address, time.Now().In(libtime.MoscowLocation()).Format("02.01.2006_15.04"), imageNumbersCache[request.DeviceNumber], extension)
	url, err := s.images.Add(ctx, name, payload, request.Payload.Size)
	if err != nil {
		return file.File{}, fmt.Errorf("failed to upload image: %w", err)
	}

	return file.File{
		Name: name,
		URL:  url,
	}, nil
}
