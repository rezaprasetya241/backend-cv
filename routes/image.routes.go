package routes

import (
	"cv/controllers"
	"github.com/gin-gonic/gin"
)

func SetupImageRoutes(router *gin.Engine) {
	imageRoutes := router.Group("/image")
	{
		//experienceRoutes.GET("/", controllers.GetEducation)
		imageRoutes.POST("/create", controllers.CreateImage)
	}
}
