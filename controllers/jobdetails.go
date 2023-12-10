package controllers

import (
	"cv/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateJob(c *gin.Context) {
	var JobData models.JobDetails

	if err := c.ShouldBindJSON(&JobData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if errCreate := models.DB.Create(&JobData).Error; errCreate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Job Details": JobData})
}
