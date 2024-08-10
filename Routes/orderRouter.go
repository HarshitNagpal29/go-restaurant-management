package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/HarshitNagpal29/go-restaurant-management/Controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orders", Controllers.GetOrders())
	incomingRoutes.GET("/orders/:order_id", Controllers.GetOrder())
	incomingRoutes.POST("/orders", Controllers.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", Controllers.UpdateOrder())
	
}