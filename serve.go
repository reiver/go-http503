package http503

import (
	"net/http"
)

func Serve(responseWriter http.ResponseWriter, retryAfter string) error {
	return ServeString(responseWriter, DefaultStatusText, retryAfter)
}
