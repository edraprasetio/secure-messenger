package main

import (
	"fmt"
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
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/api/status", StatusHandler).Methods("GET")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/protected", middlewares.ProtectedHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}

