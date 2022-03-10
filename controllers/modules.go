package controllers

import (
	"elearning/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var modules []models.Module

	if err := db.Find(&modules).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": modules})
	}
}

func Create(c *gin.Context) {
	// db := c.MustGet("db").(*gorm.DB)
	var module models.Module
	c.BindJSON(&module)
	// db.Create(&module)
	c.JSON(http.StatusOK, gin.H{"data": module})

}

func Find(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var module models.Module

	if err := db.Find(&module, c.Param("id")).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": module})
	}
}
