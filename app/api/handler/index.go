package handler

import (
	"net/http"

	"github.com/Hakitsyu/go-observability-example/internal/api"
)

func Index(writer *api.ResponseWriter, request *http.Request) {
	writer.File("api/template/index/index.html")
}
