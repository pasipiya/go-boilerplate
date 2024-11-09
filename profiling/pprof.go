package profiling

import (
	"log"
	"net/http"
)

// StartPProf starts an HTTP server for pprof profiling.
func StartPProf() {
	// Register pprof handlers
	go func() {
		log.Println("Starting pprof server on port 6060")
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			log.Fatalf("pprof server failed: %v", err)
		}
	}()
}

// Example usage of StartPProf in your main application.
