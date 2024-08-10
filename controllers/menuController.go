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

var menuCollection *mongo.Collection = Database.OpenCollection(Database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.FindOne(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching menu",
			})
		}
		var allMenus []bson.M
		if err = result.All(ctx, &allMenus); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching menu",
			})
		}
		c.JSON(http.StatusOK, allMenus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		menuId := c.Param("menu_id")
		var menu models.Menu
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menuId}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while fetching menu",
			})

		}
		c.JSON(http.StatusOK, menu)
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		validateErr := validate.Struct(menu)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validateErr.Error(),
			})
			return
		}
		menu.Menu_id = primitive.NewObjectID()
		menu.Menu_id = menu.Menu_id.Hex()
		menu.Created_at = time.Now().Format(time.RFC3339)
		menu.Updated_at = time.Now().Format(time.RFC3339)
		result, err := menuCollection.InsertOne(ctx, menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while creating menu",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.Before(start)
}
func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}
		var updateObj primitive.D
		var menu models.Menu
		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		validateErr := validate.Struct(menu)
		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validateErr.Error(),
			})
			return
		}
		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(menu.Start_Date, menu.End_Date, time.Now()) {
				msg := "Kindly provide a valid start and end date"
				c.JSON(http.StatusBadRequest, gin.H{
					"error": msg,
				})
				return
			}
		}
		updateObj = append(updateObj, bson.E{"start_date", menu.Start_Date})
		updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{"name", menu.Name})
		}
		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{"category", menu.Category})
		}
		updatedAt, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while updating menu",
			})
			return
		}
		updateObj = append(updateObj, bson.E{"updated_at", updatedAt})

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := menuCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				bson.E{"$set", updateObj},
			},
			&opt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while updating menu",
			})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
