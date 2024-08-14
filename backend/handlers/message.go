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
    var message models.Message
    json.NewDecoder(r.Body).Decode(&message)

    collection := database.GetCollection("messaging_app", "messages")
    _, err := collection.UpdateOne(
        context.TODO(),
        bson.M{"_id": message.ID},
        bson.D{{"$set", bson.D{{"content", message.Content}}}},
    )
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    objID, _ := primitive.ObjectIDFromHex(id)

    collection := database.GetCollection("messaging_app", "messages")
    _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}