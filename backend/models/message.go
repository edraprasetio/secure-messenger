package models

type Message struct {
    ID      string `bson:"_id,omitempty" json:"id,omitempty"`
    UserID  string `bson:"user_id" json:"user_id"`
    Content string `bson:"content" json:"content"`
}