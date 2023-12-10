package routes

import (
	"cv/controllers"
	"github.com/gin-gonic/gin"
)

func SetupExperienceRoutes(router *gin.Engine) {
	experienceRoutes := router.Group("/experience")
	{
		//experienceRoutes.GET("/", controllers.GetEducation)
		experienceRoutes.POST("/create", controllers.CreateExperience)
	}
}
