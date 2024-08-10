package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/HarshitNagpal29/go-restaurant-management/Controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orderItems", Controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", Controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems/order/:order_id", Controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", Controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", Controllers.UpdateOrderItem())
	
}