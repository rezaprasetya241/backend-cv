package routes

import (
	"cv/controllers"
	"github.com/gin-gonic/gin"
)

func SetupPortofolio(router *gin.Engine) {
	portofolioRoutes := router.Group("/portofolio")
	{
		//experienceRoutes.GET("/", controllers.GetEducation)
		portofolioRoutes.POST("/create", controllers.CreatePortofolio)
	}
}
