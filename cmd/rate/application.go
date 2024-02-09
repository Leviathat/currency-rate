package rate

import (
	"net/http"
)

func IncludeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/rate", RateHandler)
}
