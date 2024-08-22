package models

type Message struct {
    ID string `bson:"_id,omitempty" json:"id,omitempty"`
    SenderID string `json:"sender_id" bson:"sender_id"`
    RecipientID string `json:"recipient_id" bson:"recipient_id"`
    Content string `bson:"content" json:"content"`
    Timestamp int64 `json:"timestamp" bson:"timestamp"`
}