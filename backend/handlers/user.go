package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edraprasetio/secure-messenger/database"
	"github.com/edraprasetio/secure-messenger/models"
	"github.com/edraprasetio/secure-messenger/utils"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req RegisterRequest

	json.NewDecoder(r.Body).Decode(&req)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
	
	user := models.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	collection := database.GetCollection("secure_messenger", "users")
	_, err = collection.InsertOne(context.TODO(), user)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u User

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
	fmt.Fprint(w, tokenString)
	
}
