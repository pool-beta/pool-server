package daemon

import (
	"net/http"

	. "github.com/pool-beta/pool-server/daemon/handlers"
)

func registerRoutes(handler Handler) {
	// POOL
	http.HandleFunc("/", handler.TestHandler)
	http.HandleFunc("/test/setup", handler.TestSetup)
	http.HandleFunc("/test/reset", handler.TestReset)

	// User Routes
	http.HandleFunc("/users", handler.GetUsers)
	http.HandleFunc("/users/create", handler.CreateUser)

	// Pool Routes
	http.HandleFunc("/pool/poolid", handler.GetPool)

	http.HandleFunc("/tanks/create", handler.CreateTank)
	http.HandleFunc("/pools/create", handler.CreatePool)
	http.HandleFunc("/drains/create", handler.CreateDrain)

}
