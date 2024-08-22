package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/edraprasetio/secure-messenger/database"
	"github.com/edraprasetio/secure-messenger/models"
	"github.com/edraprasetio/secure-messenger/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	json.NewDecoder(r.Body).Decode(&req)

	collection := database.GetCollection("secure_messenger", "users")

	var existingUser models.User
	err = collection.FindOne(context.TODO(), bson.M{"username": req.Username}).Decode(&existingUser)
	if err == nil {
		// User already exists
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	} else if err != mongo.ErrNoDocuments {
		// Some other error occurred during the search
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
	
	user := models.User{
		Username: req.Username,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		Password: string(hashedPassword),
	}

	_, err = collection.InsertOne(context.TODO(), user)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Account successfully created"}`))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u models.User
	json.NewDecoder(r.Body).Decode(&u)

	collection := database.GetCollection("secure_messenger", "users")

	var user models.User
    err := collection.FindOne(context.TODO(), bson.M{"username": u.Username}).Decode(&user)
    if err != nil {
        http.Error(w, "No username found", http.StatusUnauthorized)
        return
    }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
    if err != nil {
        http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
        return
    }

	tokenString, err := utils.GenerateToken(u.Username)
	if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
        "token": tokenString,
        "user": map[string]string{
            "username":   user.Username,
            "first_name": user.FirstName,
            "last_name":  user.LastName,
            "email":      user.Email,
        },
    }

	json.NewEncoder(w).Encode(response)
	
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Get the MongoDB collection
    collection := database.GetCollection("secure_messenger", "users")

    // Prepare an empty slice to hold the user data
    var users []models.User

    // Find all user documents
    cursor, err := collection.Find(context.TODO(), bson.M{}, options.Find())
    if err != nil {
        http.Error(w, "Error fetching users", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.TODO())

    // Iterate over the cursor and decode each document into the users slice
    for cursor.Next(context.TODO()) {
        var user models.User
        cursor.Decode(&user)
        users = append(users, user)
    }

    if err := cursor.Err(); err != nil {
        http.Error(w, "Cursor error", http.StatusInternalServerError)
        return
    }

    // Return the users slice as a JSON response
    json.NewEncoder(w).Encode(users)
}

func UpdateUsername(w http.ResponseWriter, r *http.Request) {
    var req struct {
        CurrentUsername string `json:"current_username"`
        NewUsername     string `json:"new_username"`
    }

    // Parse the request body to get the current and new usernames
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Set up the MongoDB collection and filter by the current username
    collection := database.GetCollection("secure_messenger", "users")
    filter := bson.M{"username": req.CurrentUsername}

    // Set up the update operation to change the username
    update := bson.M{
        "$set": bson.M{
            "username": req.NewUsername,
        },
    }

    // Perform the update operation
    result, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        http.Error(w, "Error updating username", http.StatusInternalServerError)
        return
    }

    // Check if the username was found and updated
    if result.MatchedCount == 0 {
        http.Error(w, "Username not found", http.StatusNotFound)
        return
    }

    // Return a success response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(bson.M{"message": "Username updated successfully"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    // Parse the request body to get the username
    var req models.User
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Set up the MongoDB collection and filter by username
    collection := database.GetCollection("secure_messenger", "users")
    filter := bson.M{"username": req.Username}

    // Delete the user account from the database
    result, err := collection.DeleteOne(context.TODO(), filter)
    if err != nil || result.DeletedCount == 0 {
        http.Error(w, "User not found or delete failed", http.StatusInternalServerError)
        return
    }

    // Return a success response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(bson.M{"message": "Account deleted successfully"})
}
