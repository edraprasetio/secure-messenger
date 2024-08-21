package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/edraprasetio/secure-messenger/database"
	"github.com/edraprasetio/secure-messenger/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
    var message models.Message
    json.NewDecoder(r.Body).Decode(&message)

    collection := database.GetCollection("messaging_app", "messages")
    _, err := collection.InsertOne(context.TODO(), message)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
    collection := database.GetCollection("messaging_app", "messages")
    cursor, err := collection.Find(context.TODO(), bson.M{})
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

    collection := database.GetCollection("messaging_app", "messages")
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

    collection := database.GetCollection("messaging_app", "messages")
    filter := bson.M{"_id": objectId}

    result, err := collection.DeleteOne(context.TODO(), filter)
    if err != nil || result.DeletedCount == 0 {
        http.Error(w, "Message not found or delete failed", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(bson.M{"message": "Message deleted successfully"})
}