package http

const (
	ContentTypeHeader     = "Content-Type"
	ContentLengthHeader   = "Content-Length"
	UserAgentHeader       = "User-Agent"
	ContentEncodingHeader = "Content-Encoding"
	AcceptEncodingHeader  = "Accept-Encoding"
	AuthorizationHeader   = "Authorization"
	OriginHeader          = "Origin"
	RequestIdHeader       = "X-Request-Id"
)

const (
	ContentEncodingEmpty = ""
	ContentEncodingGzip  = "gzip"
)

const (
	ContentTypeEmpty = ""

	ContentTypeJSON = "application/json"
	ContentTypeXML  = "application/xml"
	ContentTypeHTML = "text/html"
	ContentTypeText = "text/plain"
)
