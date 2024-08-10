package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Quantity      int                `json:"quantity" bson:"quantity"`
	Unit_price    float64            `json:"unit_price" bson:"unit_price"`
	Created_at    time.Time          `json:"created_at" bson:"created_at"`
	Updated_at    time.Time          `json:"updated_at" bson:"updated_at"`
	Food_id       string             `json:"food_id" bson:"food_id"`
	Order_item_id string             `json:"order_item_id" bson:"order_item_id"`
	Order_id      string             `json:"order_id" bson:"order_id"`
}
