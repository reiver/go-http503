package http503

import (
	"net/http"
	"unsafe"
)

func ServeBytes(responseWriter http.ResponseWriter, bytes []byte, retryAfter string) error {
	var str string = unsafe.String(unsafe.SliceData(bytes), len(bytes))
	return ServeString(responseWriter, str, retryAfter)
}
