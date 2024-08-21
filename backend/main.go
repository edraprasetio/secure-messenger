package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edraprasetio/secure-messenger/database"
	"github.com/edraprasetio/secure-messenger/handlers"
	"github.com/edraprasetio/secure-messenger/middlewares"
	"github.com/gorilla/mux"
)

func main() {
	mongoURI := "mongodb://mongo:27017"
	database.ConnectMongoDB(mongoURI)
	router := mux.NewRouter()

	// Apply CORS middleware globally
	router.Use(corsMiddleware)

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/api/status", StatusHandler).Methods("GET")
	router.HandleFunc("/login", handlers.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/register", handlers.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/protected", middlewares.ProtectedHandler).Methods("GET")
	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")


	// Message Routes (protected by JWT middleware)
    router.Handle("/messages", middlewares.AuthMiddleware(http.HandlerFunc(handlers.CreateMessage))).Methods("POST")
    router.Handle("/messages", middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetMessages))).Methods("GET")
	router.Handle("/messages", middlewares.AuthMiddleware(http.HandlerFunc(handlers.UpdateMessage))).Methods("PUT")
    router.Handle("/messages", middlewares.AuthMiddleware(http.HandlerFunc(handlers.DeleteMessage))).Methods("DELETE")


	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request method and URL
		log.Printf("CORS Middleware: Handling %s request for %s", r.Method, r.URL)

		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}
