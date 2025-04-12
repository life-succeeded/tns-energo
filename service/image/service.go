package image

import (
	"fmt"
	"strings"
	"time"
	"tns-energo/cluster/analyzer"
	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
	libtime "tns-energo/lib/time"
	"tns-energo/service/file"
)

var imageNumbersCache = make(map[string]int, 32)

type Service struct {
	images          Storage
	analyzerService *analyzer.Service
}

func NewService(images Storage, analyzerService *analyzer.Service) *Service {
	return &Service{
		images:          images,
		analyzerService: analyzerService,
	}
}

func (s *Service) Upload(ctx libctx.Context, _ liblog.Logger, request UploadRequest) (file.File, error) {
	imageOptions, err := s.analyzerService.GetImageOptions(ctx, request.File)
	if err != nil {
		return file.File{}, fmt.Errorf("failed check blur score: %w", err)
	}
	if imageOptions.IsBlurred || imageOptions.HasError {
		return file.File{}, analyzer.ErrBlurredPhoto
	}

	split := strings.Split(request.File.Filename, ".")
	extension := split[len(split)-1]
	imageNumbersCache[request.DeviceNumber] = imageNumbersCache[request.DeviceNumber] + 1
	name := fmt.Sprintf("%s_%s_%d.%s", request.Address, time.Now().In(libtime.MoscowLocation()).Format("02.01.2006_15.04"), imageNumbersCache[request.DeviceNumber], extension)

	payload, err := request.File.Open()
	if err != nil {
		return file.File{}, fmt.Errorf("failed to open payload: %w", err)
	}

	url, err := s.images.Add(ctx, name, payload, request.File.Size)
	if err != nil {
		return file.File{}, fmt.Errorf("failed to upload image: %w", err)
	}

	return file.File{
		Name: name,
		URL:  url,
	}, nil
}
