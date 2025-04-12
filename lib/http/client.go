package http

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"

	libctx "tns-energo/lib/ctx"
	liblog "tns-energo/lib/log"
)

type Client interface {
	// SetVerbose включает логирование всех запросов и ответов (без тела)
	SetVerbose(v bool)
	Do(request *http.Request) (*http.Response, error)
	DoJson(ctx libctx.Context, method, url string, in, out any) (int, error)
}

type LibClient struct {
	log        liblog.Logger
	before     func(r *http.Request) error
	after      func(r *http.Response) error
	httpClient http.Client
}

func NewClient(options ...ClientOption) LibClient {
	var holder clientOptionHolder
	for _, opt := range options {
		holder = opt.apply(holder)
	}

	client := LibClient{}

	if holder.client == nil {
		client.httpClient = http.Client{}
	} else {
		client.httpClient = *holder.client
	}

	if holder.timeout > 0 {
		client.httpClient.Timeout = holder.timeout
	}

	if holder.logger != nil {
		client.log = holder.logger
	}

	if holder.transport != nil {
		client.httpClient.Transport = holder.transport
	}

	if holder.before != nil {
		client.before = holder.before
	}

	if holder.after != nil {
		client.after = holder.after
	}

	return client
}

func (c LibClient) Do(httpRequest *http.Request) (*http.Response, error) {
	if c.before != nil {
		if err := c.before(httpRequest); err != nil {
			return nil, fmt.Errorf("не удалось выполнить дополнительную подготовку запроса: %v", err)
		}
	}

	response, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	count := response.ContentLength
	if count < 0 && response.Header != nil {
		count, _ = strconv.ParseInt(response.Header.Get(ContentLengthHeader), 10, 64)
	}

	if c.after != nil {
		if err = c.after(response); err != nil {
			return nil, fmt.Errorf("не удалось выполнить дополнительную обработку ответа: %v", err)
		}
	}

	return response, nil
}

func (c LibClient) DoJson(ctx libctx.Context, method, url string, in, out any) (int, error) {
	rq, err := NewRequest(ctx, method, url, nil)
	if err != nil {
		return 0, err
	}

	if err = WriteRequestJson(rq, in); err != nil {
		return 0, err
	}

	rs, err := c.Do(rq)
	if err != nil {
		return 0, err
	}

	if err = ReadResponseJson(rs, out); err != nil {
		return rs.StatusCode, err
	}

	return rs.StatusCode, nil
}

type FormDataField struct {
	Name  string
	Value string
}

type FormDataFile struct {
	FieldName string
	FileName  string
	Base64    string
}

// SendFormData send multipart/form-data request
func (c LibClient) SendFormData(ctx libctx.Context, url string, fields []FormDataField, files []FormDataFile, out any) (int, error) {
	var (
		body   = &bytes.Buffer{}
		writer = multipart.NewWriter(body)
	)

	for _, field := range fields {
		err := writer.WriteField(field.Name, field.Value)
		if err != nil {
			return http.StatusBadRequest, fmt.Errorf("can't write files: %w", err)
		}
	}

	for _, file := range files {
		decodedData, err := base64.StdEncoding.DecodeString(file.Base64)
		if err != nil {
			return http.StatusBadRequest, fmt.Errorf("can't decode files: %w", err)
		}

		part, err := writer.CreateFormFile(file.FieldName, file.FileName)
		if err != nil {
			return http.StatusBadRequest, fmt.Errorf("can't create form file: %w", err)
		}

		_, err = io.Copy(part, bytes.NewReader(decodedData))
		if err != nil {
			return http.StatusBadRequest, fmt.Errorf("can't copy file: %w", err)
		}
	}

	if err := writer.Close(); err != nil {
		return http.StatusBadRequest, fmt.Errorf("can't close multipart-writer: %w", err)
	}

	req, err := NewRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("can't create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	rs, err := c.httpClient.Do(req)
	if err != nil {
		return http.StatusBadRequest, fmt.Errorf("can't create request: %w; rs status code: %d", err, rs.StatusCode)
	}

	if err = ReadResponseJson(rs, out); err != nil {
		return rs.StatusCode, err
	}

	return rs.StatusCode, nil
}
