package api

import (
	"net/http"
)

var serverMux = http.NewServeMux()

func StartServer() {
	InitializeAPI()

	http.ListenAndServe(":8080", serverMux)
}
