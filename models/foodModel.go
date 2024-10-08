package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Price      float64            `json:"price" bson:"price"`
	Food_image string             `json:"food_image" bson:"food_image"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
	Food_id    string             `json:"food_id" bson:"food_id"`
	Menu_id    *string            `json:"menu_id" bson:"menu_id"`
}
