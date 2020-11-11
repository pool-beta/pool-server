package daemon

import (
	"log"
	"net/http"

	. "github.com/pool-beta/pool-server/daemon/handlers"
)

func Run() {
	
	// Start Handlers
	handler, err := NewHandler()
	if err != nil {
		log.Fatal("Error in setting up handler")
	}

	// Register Routes
	registerRoutes(handler)

	// Start Http Server
    log.Println("POOL Backend Server has be started. Listening on http://localhost:8000.")
	log.Fatal(http.ListenAndServe(":8000", nil))
}