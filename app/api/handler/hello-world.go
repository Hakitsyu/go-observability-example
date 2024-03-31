package handler

import "net/http"

func HelloWorld(writer http.ResponseWriter, request *http.Request) {
	response := "<h1>Hello World</h1>"

	writer.Write([]byte(response))
}
