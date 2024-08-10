package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	First_Name    string             `json:"first_name" bson:"first_name"`
	Last_Name     string             `json:"last_name" bson:"last_name"`
	Password      string             `json:"password" bson:"password"`
	Email         string             `json:"email" bson:"email"`
	Avatar        string             `json:"avatar" bson:"avatar"`
	Phone         string             `json:"phone" bson:"phone"`
	Token         string             `json:"token" bson:"token"`
	Refresh_Token string             `json:"refresh_token" bson:"refresh_token"`
	Created_at    time.Time          `json:"created_at" bson:"created_at"`
	Updated_at    time.Time          `json:"updated_at" bson:"updated_at"`
	User_id       string             `json:"user_id" bson:"user_id"`
}
