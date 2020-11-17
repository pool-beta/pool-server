package daemon

import (
	"net/http"

	. "github.com/pool-beta/pool-server/daemon/handlers"
)

func registerRoutes(handler Handler) {
	// POOL
	http.HandleFunc("/", handler.TestHandler)

	// User Routes
	http.HandleFunc("/users", handler.GetUsers)
	http.HandleFunc("/users/create", handler.CreateUser)

	// Pool Routes
	http.HandleFunc("/user/tanks/create", handler.CreateTank)
}
