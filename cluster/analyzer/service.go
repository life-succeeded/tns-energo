package analyzer

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/sunshineOfficial/golib/goctx"
	"github.com/sunshineOfficial/golib/gohttp"
)

type Service struct {
	client  gohttp.Client
	baseUrl string
}

func NewService(client gohttp.Client, baseUrl string) *Service {
	return &Service{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (s *Service) GetImageOptions(ctx goctx.Context, file *multipart.FileHeader) (ImageQualityResult, error) {
	payload, err := file.Open()
	if err != nil {
		return ImageQualityResult{}, fmt.Errorf("failed to open payload: %w", err)
	}

	var files = []gohttp.FormDataFile{{
		Payload:    payload,
		MIMEHeader: file.Header,
	}}

	var response ImageQualityResult
	status, err := s.client.SendFormData(ctx, fmt.Sprintf("%s/process-image", s.baseUrl), nil, files, &response)
	if err != nil {
		return ImageQualityResult{}, fmt.Errorf("failed to upload image: %w", err)
	}
	if status != http.StatusOK {
		return ImageQualityResult{}, fmt.Errorf("bad response, status: %d", status)
	}

	return response, nil
}
