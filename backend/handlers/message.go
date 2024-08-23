package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/edraprasetio/secure-messenger/database"
	"github.com/edraprasetio/secure-messenger/models"
	"github.com/edraprasetio/secure-messenger/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
    // Get the token from the request header (already verified by the middleware)
    authHeader := r.Header.Get("Authorization")
    tokenString := strings.Split(authHeader, "Bearer ")[1]

    // Extract the username from the token
    senderUsername, err := utils.GetUsernameFromToken(tokenString)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var message models.Message
    json.NewDecoder(r.Body).Decode(&message)
    
    // Set the Sender to the current user's username
    message.Sender = senderUsername

    message.Timestamp = time.Now().Unix()

    // Insert the message into the database
    messageCollection := database.GetCollection("messaging_app", "messages")
    _, err = messageCollection.InsertOne(context.TODO(), message)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
    // Get the token from the request header (already verified by the middleware)
    authHeader := r.Header.Get("Authorization")
    tokenString := strings.Split(authHeader, "Bearer ")[1]

    // Extract the username from the token
    username, err := utils.GetUsernameFromToken(tokenString)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    // Find messages sent by or to the current user
    messageCollection := database.GetCollection("messaging_app", "messages")
    filter := bson.M{
        "$or": []bson.M{
            {"sender": username},
            {"recipient": username},
        },
    }
    
    cursor, err := messageCollection.Find(context.TODO(), filter)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    var messages []models.Message
    if err := cursor.All(context.TODO(), &messages); err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(messages)
}


func UpdateMessage(w http.ResponseWriter, r *http.Request) {
    var req models.Message
    json.NewDecoder(r.Body).Decode(&req)

    objectId, err := primitive.ObjectIDFromHex(req.ID)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    collection := database.GetCollection("secure_messenger", "messages")
    filter := bson.M{"_id": objectId}

    update := bson.M{
        "$set": bson.M{
            "content": req.Content,
        },
    }

    result, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil || result.MatchedCount == 0 {
        http.Error(w, "Message not found or update failed", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(bson.M{"message": "Message updated successfully"})
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
    var req models.Message
    json.NewDecoder(r.Body).Decode(&req)

    objectId, err := primitive.ObjectIDFromHex(req.ID)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    collection := database.GetCollection("secure_messenger", "messages")
    filter := bson.M{"_id": objectId}

    result, err := collection.DeleteOne(context.TODO(), filter)
    if err != nil || result.DeletedCount == 0 {
        http.Error(w, "Message not found or delete failed", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(bson.M{"message": "Message deleted successfully"})
}