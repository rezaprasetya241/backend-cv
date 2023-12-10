package routes

import (
	"cv/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.POST("/create", controllers.CreateUser)
	}
}
