package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/HarshitNagpal29/go-restaurant-management/Controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/menus", Controllers.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", Controllers.GetMenu())
	incomingRoutes.POST("/menus", Controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:menu_id", Controllers.UpdateMenu())	
}