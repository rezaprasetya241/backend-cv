package controllers

import (
	"cv/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

func GetUser(c *gin.Context) {
	var UserDetails models.User
	id := c.Param("id")
	//Preload(clause.Associations).Preload("Experiences."+clause.Associations).Preload("Portofolio."+clause.Associations)
	if err := models.DB.Preload(clause.Associations).Preload("Experiences."+clause.Associations).Preload("Portofolio."+clause.Associations).First(&UserDetails, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": UserDetails})
}

func CreateUser(c *gin.Context) {
	var UserData models.User

	if err := c.ShouldBindJSON(&UserData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if errCreate := models.DB.Create(&UserData).Error; errCreate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": UserData})
}
