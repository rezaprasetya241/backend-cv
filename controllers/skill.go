package controllers

import (
	"cv/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSkill(c *gin.Context) {
	var Skill models.Skill

	if err := c.ShouldBindJSON(&Skill); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if errCreate := models.DB.Create(&Skill).Error; errCreate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errCreate.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Skill": Skill})
}
