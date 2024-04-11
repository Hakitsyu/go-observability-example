package api

import (
	"log"
	"net/http"
	"os"
)

type ResponseWriter struct {
	writer http.ResponseWriter
}

func (writer *ResponseWriter) String(content string) {
	writer.writer.Write([]byte(content))
}

func (writer *ResponseWriter) Error() {
	writer.writer.Write([]byte("Error..."))
}

func (writer *ResponseWriter) File(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		writer.Error()
		return
	}

	writer.writer.Write(content)
}

func NewResponseWriter(writer http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{writer}
}
