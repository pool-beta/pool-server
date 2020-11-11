package daemon

import (
	"log"
	"net/http"

	. "github.com/pool-beta/pool-server/daemon/handlers"
)

const PORT = ":8000"

func Run() {
	
	// Start Handlers
	// Should only be called once (singleton)
	handler, err := NewHandler()
	if err != nil {
		log.Fatal("Error in setting up handler")
	}

	// Register Routes
	registerRoutes(handler)

	// Start Http Server
    log.Printf("POOL Backend Server has be started. Listening on http://localhost%v.", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}