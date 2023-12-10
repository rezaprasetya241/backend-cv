package routes

import (
	"cv/controllers"
	"github.com/gin-gonic/gin"
)

func SetupEducationRoutes(router *gin.Engine) {
	educationRoutes := router.Group("/education")
	{
		educationRoutes.GET("/", controllers.GetEducation)
		educationRoutes.POST("/create", controllers.CreateEducation)
	}
}
