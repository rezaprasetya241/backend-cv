package controllers

import (
	"cv/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePortofolio(c *gin.Context) {
	var Portofolio models.Portofolio

	if err := c.ShouldBindJSON(&Portofolio); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if errCreate := models.DB.Create(&Portofolio).Error; errCreate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Portofolio": Portofolio})
}
