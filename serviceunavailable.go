package http503

import (
	"net/http"
)

func ServiceUnavailable(responseWriter http.ResponseWriter, request *http.Request, retryAfter string) {
	Serve(responseWriter, retryAfter)
}
