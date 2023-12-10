package controllers

import (
	"cv/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateImage(c *gin.Context) {
	var Image models.ImageUrl

	if err := c.ShouldBindJSON(&Image); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if errCreate := models.DB.Create(&Image).Error; errCreate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Image": Image})
}
