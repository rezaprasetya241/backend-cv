package routes

import (
	"cv/controllers"
	"github.com/gin-gonic/gin"
)

func SetupCertificateRoutes(router *gin.Engine) {
	certificateRoutes := router.Group("/certiificate")
	{
		//certificateRoutes.GET("/", controllers.GetEducation)
		certificateRoutes.POST("/create", controllers.CreateCertificate)
	}
}
