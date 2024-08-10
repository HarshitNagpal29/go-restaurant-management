package Controllers

import (
	"context"
	"math"
	"net/http"
	"time"

	"github.com/HarshitNagpal29/go-restaurant-management/Database"
	"github.com/HarshitNagpal29/go-restaurant-management/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection *mongo.Collection = Database.OpenCollection(Database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var allFoods bson.M
		result, err := foodCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching foods",
			})

		}
		if err = result.All(ctx, &allFoods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching foods",
			})

		}
		c.JSON(http.StatusOK, allFoods)
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching food",
			})
			c.JSON(http.StatusOK, food)

		}

	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

		}
		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationErr.Error(),
			})
			return
		}
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Menu not found",
			})
			return
		}
		createdTime := time.Now().Format(time.RFC3339)
		food.Created_at, _ = time.Parse(time.RFC3339, createdTime)
		updatedTime, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at = updatedTime
		food.Food_id = primitive.NewObjectID().Hex()
		food.Food_id = food.ID.Hex()
		num := toFixed(food.Price, 2)
		food.Price = num

		result, insertErr := foodCollection.InsertOne(ctx, food)
		if insertErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while creating food",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu
		var updateObj primitive.D
		defer cancel()
		foodId := c.Param("food_id")
		filter := bson.M{"food_id": foodId}

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		validateErr := validate.Struct(&food)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validateErr.Error(),
			})

			err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Menu not found",
				})
				return
			}

			if food.Name != "" {
				updateObj = append(updateObj, bson.E{"name", food.Name})
			}

			if food.Price != 0 {
				updateObj = append(updateObj, bson.E{"price", food.Price})
			}

			if food.Food_image != "" {
				updateObj = append(updateObj, bson.E{"food_image", food.Food_image})
			}

			if food.Menu_id != nil {
				updateObj = append(updateObj, bson.E{"menu_id", food.Menu_id})
			}

			food.Updated_at = time.Now().Format(time.RFC3339)
			updateObj = append(updateObj, bson.E{"updated_at", food.Updated_at})

			upsert := true
			opt := options.UpdateOptions{
				Upsert: &upsert,
			}
			result, err := foodCollection.UpdateOne(
				ctx,
				filter,
				bson.D{
					bson.E{"$set", updateObj},
				},
				&opt,
			)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Error while updating food",
				})
				return
			}
			defer cancel()
			c.JSON(http.StatusOK, result)

		}
	}
}
