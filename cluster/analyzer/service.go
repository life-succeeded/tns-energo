package analyzer

import (
	"fmt"
	"mime/multipart"
	"net/http"
	libctx "tns-energo/lib/ctx"
	libhttp "tns-energo/lib/http"
)

const fileField = "file"

type Service struct {
	client  libhttp.Client
	baseUrl string
}

func NewService(client libhttp.Client, baseUrl string) *Service {
	return &Service{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (s *Service) GetImageOptions(ctx libctx.Context, fileName string, payload multipart.File) (ImageQualityResult, error) {
	var files = []libhttp.FormDataFile{{
		FieldName: fileField,
		FileName:  fileName,
		Payload:   payload,
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
