package models

type Message struct {
    ID          string `bson:"_id,omitempty" json:"id,omitempty"`
    Sender      string `json:"sender" bson:"sender"`
    Recipient   string `json:"recipient" bson:"recipient"`
    Content     string `bson:"content" json:"content"`
    Timestamp   int64 `json:"timestamp" bson:"timestamp"`
}