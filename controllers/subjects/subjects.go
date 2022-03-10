package controllers

import (
	"elearning/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var subjects []models.Subject

	if err := db.Find(&subjects).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": subjects})
	}
}

func Find(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var subject models.Subject

	if err := db.Find(&subject, c.Param("id")).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": subject})
	}
}
