package main

import (
	"cv/initializers"
	"cv/models"
	"cv/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
	models.ConnectDatabase()
}

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Reza Prasetya",
			"bio":  "Belajar backend go",
		})
	})

	routes.SetupEducationRoutes(router)
	routes.SetupUserRoutes(router)
	routes.SetupCertificateRoutes(router)
	routes.SetupSkillRoutes(router)
	routes.SetupExperienceRoutes(router)
	routes.SetupJobDetailsRoutes(router)
	routes.SetupPortofolio(router)
	routes.SetupImageRoutes(router)

	log.Fatal(router.Run(":" + port))
}
