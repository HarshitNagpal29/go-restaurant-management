package Controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/HarshitNagpal29/go-restaurant-management/Database"
	"github.com/HarshitNagpal29/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = Database.OpenCollection(Database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var allOrders bson.M
		result, err := orderCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching orders",
			})
		}
		if err = result.All(ctx, &allOrders); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching orders",
			})
		}
		c.JSON(http.StatusOK, allOrders)

	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("order_id")
		var order models.Order
		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}.Decode(&order))
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching order",
			})
		}
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var order models.Order
		var table models.Table

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error while binding order",
			})
			return
		}
		validateErr := validate.Struct(&order)
		if validateErr != nil {
			c.BindJSON(http.StatusBadRequest, gin.H{
				"error": "Error while validating order",
			})
			return
		}
		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error while fetching table",
				})
				return
			}
		}
		createdTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Created_at = createdTime
		updatedTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at = updatedTime
		order.Order_id = primitive.NewObjectID()
		orderID, _ := primitive.ObjectIDFromHex(order.Order_id)
		order.Order_id = orderID.Hex()
		result, insertErr := orderCollection.InsertOne(ctx, order)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while creating order",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("order_id")
		var order models.Order
		var table models.Table
		var updateObj primitive.D

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		validateErr := validate.Struct(&order)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validateErr.Error(),
			})
		}
		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error while fetching table",
				})
			}
			updateObj = append(updateObj, bson.E{"table_id", order.Table_id})
		}
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})
		filter := bson.M{"order_id": orderId}
		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				bson.E{"$set", updateObj},
			},
			&opt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while updating order",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func OrderItemOrderCreator(order models.Order) string {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Order_id = primitive.NewObjectID()
	orderID, _ := primitive.ObjectIDFromHex(order.Order_id)
	order.Order_id = orderID.Hex()
	orderCollection.InsertOne(ctx, order)

	return order.Order_id
}
