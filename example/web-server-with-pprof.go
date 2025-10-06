package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Enable pprof endpoints
	"sync"
	"time"
)

// Example web server with intentional goroutine leaks for testing LeakFinder

var (
	// Simulate some global state that might cause leaks
	activeConnections = make(map[string]*Connection)
	connectionsMutex  sync.RWMutex
)

type Connection struct {
	ID        string
	StartTime time.Time
	Done      chan bool
}

// Intentionally leaky function - creates goroutines that never exit
func leakyWorker(id string) {
	conn := &Connection{
		ID:        id,
		StartTime: time.Now(),
		Done:      make(chan bool),
	}

	connectionsMutex.Lock()
	activeConnections[id] = conn
	connectionsMutex.Unlock()

	// This goroutine will never exit - simulates a leak!
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// Simulate some work
				_ = fmt.Sprintf("Working on connection %s", id)
			case <-conn.Done:
				// This will never be called in our leaky example
				return
			}
		}
	}()
}

// Handler that creates leaky goroutines
func leakyHandler(w http.ResponseWriter, r *http.Request) {
	connectionID := fmt.Sprintf("conn_%d", time.Now().UnixNano())

	// Create a leaky worker
	leakyWorker(connectionID)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "created", "connection_id": "%s", "active_connections": %d}`,
		connectionID, len(activeConnections))
}

// Handler to show current stats
func statsHandler(w http.ResponseWriter, r *http.Request) {
	connectionsMutex.RLock()
	count := len(activeConnections)
	connectionsMutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"active_connections": %d, "message": "Use LeakFinder to detect the goroutine leaks!"}`, count)
}

// Handler for health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "healthy", "message": "Server is running with pprof enabled on :6060"}`)
}

func main() {
	// Set up HTTP routes
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/leak", leakyHandler)
	http.HandleFunc("/stats", statsHandler)

	// Start pprof server on port 6060 (default for LeakFinder)
	go func() {
		log.Println("🔍 pprof server starting on :6060")
		log.Println("   Available endpoints:")
		log.Println("   - http://localhost:6060/debug/pprof/")
		log.Println("   - http://localhost:6060/debug/pprof/goroutine?debug=2")
		log.Println("   - http://localhost:6060/debug/pprof/heap")
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Printf("pprof server error: %v", err)
		}
	}()

	// Start main web server on port 8080
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	log.Println("🚀 Example web server starting on :8080")
	log.Println("   Available endpoints:")
	log.Println("   - http://localhost:8080/health - Health check")
	log.Println("   - http://localhost:8080/leak - Create goroutine leak")
	log.Println("   - http://localhost:8080/stats - Show current stats")
	log.Println("")
	log.Println("💡 To test with LeakFinder:")
	log.Println("   1. Start this server: go run examples/web-server-with-pprof.go")
	log.Println("   2. Create some leaks: curl http://localhost:8080/leak")
	log.Println("   3. Run LeakFinder: ./leakfinder check")
	log.Println("   4. Analyze leaks: ./leakfinder analyze")
	log.Println("   5. Monitor real-time: ./leakfinder monitor")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
