package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"goLangFirst/internal/config"
	"goLangFirst/internal/http/handlers/student"
	"goLangFirst/internal/storage/sqlite"
)

func main() {
	// Load config
	cfg := config.MustLoad()
	// Database connection
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Storage initialized", slog.String("path", cfg.StoragePath), slog.String("env", cfg.Env))
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))

	// Setup server
	server := &http.Server{
		Addr:    cfg.HTTPServer.Addr, // Use the correct field
		Handler: router,
	}

	fmt.Printf("Server starting at %s\n", cfg.HTTPServer.Addr)
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
	<-done
	slog.Info("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("Server shut down seccessfully")

}
