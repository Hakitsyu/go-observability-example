package middleware

import "net/http"

type StatusCodeRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (recorder *StatusCodeRecorder) WriteHeader(statusCode int) {
	recorder.statusCode = statusCode
	recorder.ResponseWriter.WriteHeader(statusCode)
}

func NewStatusCodeRecorder(writer http.ResponseWriter) StatusCodeRecorder {
	return StatusCodeRecorder{writer, 200}
}
