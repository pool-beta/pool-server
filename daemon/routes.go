package daemon

import (
	"net/http"

	. "github.com/pool-beta/pool-server/daemon/handlers"
)

func registerRoutes(handler Handler) {
	http.HandleFunc("/test", handler.TestHandler)
}