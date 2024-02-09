// api/application.go

package api

import (
	"currency-rate/cmd/rate"
)

func InitializeAPI() {
	rate.IncludeRoutes(serverMux)
}
