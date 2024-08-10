package routes

import (
	"github.com/HarshitNagpal29/go-restaurant-management/Controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", Controllers.GetFoods())
	incomingRoutes.GET("/foods/:food_id", Controllers.GetFood())
	incomingRoutes.POST("/foods", Controllers.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", Controllers.UpdateFood())
}
