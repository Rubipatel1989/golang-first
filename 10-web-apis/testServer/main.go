package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Response represents a standard API response
type Response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
	Status  string `json:"status"`
}

// User represents a user entity
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Pawan Kumar", Email: "pawan@example.com"},
	{ID: 2, Name: "John Doe", Email: "john@example.com"},
}

// Health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Server is running!",
		Time:    time.Now().Format(time.RFC3339),
		Status:  "healthy",
	}
	json.NewEncoder(w).Encode(response)
}

// Get all users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get user by ID
func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	// Simple ID matching (in real app, convert to int and search)
	fmt.Fprintf(w, `{"message": "Get user by ID: %s", "time": "%s"}`, id, time.Now().Format(time.RFC3339))
}

// Root endpoint
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Welcome to Go Learning Project API!",
		Time:    time.Now().Format(time.RFC3339),
		Status:  "active",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	port := ":8080"
	
	// Register routes
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/users", getUsersHandler)
	http.HandleFunc("/api/user", getUserByIDHandler)
	
	fmt.Printf("🚀 Server starting on port %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  http://localhost:8080/")
	fmt.Println("  GET  http://localhost:8080/health")
	fmt.Println("  GET  http://localhost:8080/api/users")
	fmt.Println("  GET  http://localhost:8080/api/user?id=1")
	
	log.Fatal(http.ListenAndServe(port, nil))
}
