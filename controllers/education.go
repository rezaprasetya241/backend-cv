package controllers

import (
	"cv/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEducation(c *gin.Context) {
	var educations []models.Education

	models.DB.Find(&educations)
	c.JSON(http.StatusOK, gin.H{
		"data": educations,
	})
}

func CreateEducation(c *gin.Context) {
	var Educations models.Education

	if err := c.ShouldBindJSON(&Educations); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if errCreate := models.DB.Create(&Educations).Error; errCreate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Educations": Educations})
}
