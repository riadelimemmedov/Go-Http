package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//! Define server configuration
	const port = ":8080"

	// ! Create a multiplexer to handle different routes
	mux := http.NewServeMux()

	//!Register route handlers
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api/data", dataHandler)

	server := &http.Server{
		Addr:    port,
		Handler: mux,

		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second, //!Keep open connections for 60 seconds
		ReadHeaderTimeout: 5 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1 MB,
	}

	go func() {
		fmt.Printf("Server is starting on port %s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not start server: %v", err)
		}
	}()
	setupGracefulShutdown(server)
}

// ! healthHandler provides a health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	health := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"uptime":    "server is running",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(health)
}

// ! Home handler function to serve the root path main page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	//!Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//!Create response data
	response := map[string]interface{}{
		"message":    "Welcome to the Go HTTP Server!",
		"timestamp":  time.Now().UTC().Format(time.RFC3339),
		"method":     r.Method,
		"path":       r.URL.Path,
		"userAgent":  r.Header.Get("User-Agent"),
		"remoteAddr": r.RemoteAddr,
	}

	//!Encode and send JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// ! dataHandler is a simulates an API endpoint with some processing time
func dataHandler(w http.ResponseWriter, r *http.Request) {
	//! Simulate processing time
	time.Sleep(100 * time.Millisecond)

	w.Header().Set("Content-Type", "application/json")

	data := map[string]interface{}{
		"data": map[string]interface{}{
			"id":   1,
			"name": "Sample Data",
			"info": "This is a sample response from the data endpoint.",
		},
		"count":     1,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// ! setupGracefulShutdown handles clean server shutdown on interrupt signals
func setupGracefulShutdown(server *http.Server) {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-sigChan
	fmt.Printf("\n🛑 Received signal: %v\n", sig)
	fmt.Println("🔄 Initiating graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("❌ Server forced to shutdown: %v", err)
	} else {
		fmt.Println("✅ Server gracefully stopped")
	}
}
