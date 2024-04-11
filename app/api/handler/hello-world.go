package handler

import (
	"net/http"

	"github.com/Hakitsyu/go-observability-example/internal/api"
)

func HelloWorld(writer *api.ResponseWriter, request *http.Request) {
	writer.String("<h1>Hello World</h1>")
}
