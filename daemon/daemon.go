package daemon

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/pool-beta/pool-server/daemon/handlers"
)

func Run(port string) {
	
	// Start Handlers
	// Should only be called once (singleton)
	handler, err := NewHandler()
	if err != nil {
		log.Fatal("Error in setting up handler")
	}

	// Register Routes
	registerRoutes(handler)

	// Start Http Server
    log.Printf("POOL Backend Server has been started. Listening on http://localhost:%v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}