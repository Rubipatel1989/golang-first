package main

import (
	"fmt"
	"log"
	"net/http"

	"goLangFirst/internal/config"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// Setup server
	server := &http.Server{
		Addr:    cfg.HTTPServer.Addr, // Use the correct field
		Handler: router,
	}

	fmt.Printf("Server starting at %s\n", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
