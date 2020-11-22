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

	http.HandleFunc("/user/tanks/create", handler.CreateTank)
	http.HandleFunc("/user/pools/create", handler.CreatePool)
	http.HandleFunc("/user/drains/create", handler.CreateDrain)

}
