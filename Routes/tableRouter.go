package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/HarshitNagpal29/go-restaurant-management/Controllers"
)

func TableRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/tables", Controllers.GetTables())
	incomingRoutes.GET("/tables/:table_id", Controllers.GetTable())
	incomingRoutes.POST("/tables", Controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", Controllers.UpdateTable())
	
}
