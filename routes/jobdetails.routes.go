package routes

import (
	"cv/controllers"
	"github.com/gin-gonic/gin"
)

func SetupJobDetailsRoutes(router *gin.Engine) {
	jobRoutes := router.Group("/jobdetails")
	{
		//experienceRoutes.GET("/", controllers.GetEducation)
		jobRoutes.POST("/create", controllers.CreateJob)
	}
}
