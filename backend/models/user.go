package models

type User struct {
    ID       string `bson:"_id,omitempty" json:"id,omitempty"`
    Username string `bson:"username" json:"username"`
    Email string `bson:"email" json:"email"`
    Password string `bson:"password" json:"password"`
}