package analyzer

import (
	"fmt"
	statusCodes "net/http"
	libctx "tns-energo/lib/ctx"
	"tns-energo/lib/http"
)

const fileFiled = "file"

type Service struct {
	client  *http.LibClient
	baseUrl string
}

func NewService(client *http.LibClient, baseUrl string) *Service {
	return &Service{
		client:  client,
		baseUrl: baseUrl,
	}
}

func (s *Service) GetImageOptions(ctx libctx.Context, base64Payload string) (ImageQualityResult, error) {
	var response ImageQualityResult

	var files = []http.FormDataFile{{
		FieldName: fileFiled,
		Base64:    base64Payload,
		FileName:  "image.jpg",
	}}

	statusCode, err := s.client.SendFormData(ctx, fmt.Sprintf("%s/process-image", s.baseUrl), nil, files, &response)
	if err != nil {
		return response, err
	}
	if statusCode != statusCodes.StatusOK {
		return response, fmt.Errorf("bad response, status: %d", statusCode)
	}

	return response, nil
}
