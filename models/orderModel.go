package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Order_Date time.Time          `json:"order_date" bson:"order_date"`
	Created_at time.Time          `json:"creation_date" bson:"creation_date"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
	Order_id   string             `json:"order_id" bson:"order_id"`
	Table_id   *string            `json:"table_id" bson:"table_id"`
}
