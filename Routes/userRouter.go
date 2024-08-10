package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/HarshitNagpal29/go-restaurant-management/Controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", Controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", Controllers.GetUser())
	incomingRoutes.POST("/users/signup", Controllers.SignUp())
	incomingRoutes.POST("/users/login", Controllers.Login())
}