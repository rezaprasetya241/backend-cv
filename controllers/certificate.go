package controllers

import (
	"cv/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func GetSkill(c *gin.Context) {
//	//var skill []models.Education
//	var skill models.Education
//
//	id := c.Param("id")
//	models.DB.Find(&educations)
//	c.JSON(http.StatusOK, gin.H{
//		"data": educations,
//	})
//}

func CreateCertificate(c *gin.Context) {
	var Certificate models.Certificate

	if err := c.ShouldBindJSON(&Certificate); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if errCreate := models.DB.Create(&Certificate).Error; errCreate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Certificate": Certificate})
}
