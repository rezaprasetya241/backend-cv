package controllers

import (
	"cv/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateExperience(c *gin.Context) {
	var Experience models.Experience

	startString := Experience.StartDate.Format("2006-01-02")
	startDate, _ := time.Parse("2006-01-02 3:04PM", startString)
	endString := Experience.EndDate.Format("2006-01-02")
	endDate, _ := time.Parse("2006-01-02 3:04PM", endString)

	Experience.StartDate = startDate
	Experience.EndDate = endDate

	fmt.Println("startDate: ", Experience.StartDate)

	fmt.Println("startDate: ", Experience.StartDate)
	if err := c.ShouldBindJSON(&Experience); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if errCreate := models.DB.Create(&Experience).Error; errCreate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Experience": Experience})
}
