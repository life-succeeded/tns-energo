package http

import (
	"fmt"
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

type SwClient struct {
	log        liblog.Logger
	before     func(r *http.Request) error
	after      func(r *http.Response) error
	httpClient http.Client
}

func NewClient(options ...ClientOption) SwClient {
	var holder clientOptionHolder
	for _, opt := range options {
		holder = opt.apply(holder)
	}

	client := SwClient{}

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

func (c SwClient) Do(httpRequest *http.Request) (*http.Response, error) {
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

func (c SwClient) DoJson(ctx libctx.Context, method, url string, in, out any) (int, error) {
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
