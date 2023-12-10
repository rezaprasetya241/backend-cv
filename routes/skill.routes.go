package routes

import (
	"cv/controllers"
	"github.com/gin-gonic/gin"
)

func SetupSkillRoutes(router *gin.Engine) {
	skillRoutes := router.Group("/skill")
	{
		//skillRoutes.GET("/", controllers.)
		skillRoutes.POST("/create", controllers.CreateSkill)
	}
}
