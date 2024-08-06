package main

import (
	// "context"
	"fmt"
	// "log"
	"net/http"
	// "time"

	"github.com/edraprasetio/secure-messenger/database"
	"github.com/edraprasetio/secure-messenger/handlers"
	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// var client *mongo.Client

func main() {
	mongoURI := "mongodb://mongo:27017"
	database.ConnectMongoDB(mongoURI)
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/api/status", StatusHandler).Methods("GET")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", handlers.ProtectedHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}

// func connectMongoDB() (*mongo.Client, error) {
//     mongoURI := "mongodb://mongo:27017"
//     clientOptions := options.Client().ApplyURI(mongoURI)

//     client, err := mongo.Connect(context.TODO(), clientOptions)
//     if err != nil {
//         return nil, err
//     }

//     ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//     defer cancel()

//     err = client.Ping(ctx, nil)
//     if err != nil {
//         return nil, err
//     }

//     fmt.Println("Connected to MongoDB!")
//     return client, nil
// }

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK")
}

// func TestMongoHandler(w http.ResponseWriter, r *http.Request) {
//     err := client.Ping(context.TODO(), nil)
//     if err != nil {
//         http.Error(w, "Failed to connect to MongoDB", http.StatusInternalServerError)
//         return
//     }
//     fmt.Fprintf(w, "Successfully connected to MongoDB!")
// }
