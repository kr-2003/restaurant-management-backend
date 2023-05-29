package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kr-2003/restaurant-management/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/login", controller.Login())
	incomingRoutes.POST("/signup", controller.SignUp())
}
