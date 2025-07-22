package http503

import (
	"net/http"
	"io"
)

func ServeString(responseWriter http.ResponseWriter, value string, retryAfter string) error {
	if nil == responseWriter {
		return ErrNilHTTPResponseWriter
	}

	responseWriter.Header().Add("Retry-After", retryAfter)

	responseWriter.WriteHeader(StatusCode)

	if "" != value {
		io.WriteString(responseWriter, value)
	}

	return nil
}
