package http

import (
	"net/http"
	"time"

	liblog "tns-energo/lib/log"
)

type clientOptionHolder struct {
	client    *http.Client
	transport http.RoundTripper
	logger    liblog.Logger
	before    func(*http.Request) error
	after     func(*http.Response) error
	verbose   bool
	timeout   time.Duration
}

type ClientOption interface {
	apply(options clientOptionHolder) clientOptionHolder
}

type clientOptionFunc func(o clientOptionHolder) clientOptionHolder

func (f clientOptionFunc) apply(o clientOptionHolder) clientOptionHolder {
	return f(o)
}

func WithLogger(logger liblog.Logger) ClientOption {
	return clientOptionFunc(func(o clientOptionHolder) clientOptionHolder {
		o.logger = logger
		o.verbose = true
		return o
	})
}

func WithClient(client *http.Client) ClientOption {
	return clientOptionFunc(func(o clientOptionHolder) clientOptionHolder {
		o.client = client
		return o
	})
}

func WithTransport(transport http.RoundTripper) ClientOption {
	return clientOptionFunc(func(o clientOptionHolder) clientOptionHolder {
		o.transport = transport
		return o
	})
}

func WithTimeout(timeout time.Duration) ClientOption {
	return clientOptionFunc(func(o clientOptionHolder) clientOptionHolder {
		o.timeout = timeout
		return o
	})
}

func WithBefore(f func(r *http.Request) error) ClientOption {
	return clientOptionFunc(func(o clientOptionHolder) clientOptionHolder {
		o.before = f
		return o
	})
}

func WithAfter(f func(r *http.Response) error) ClientOption {
	return clientOptionFunc(func(o clientOptionHolder) clientOptionHolder {
		o.after = f
		return o
	})
}
