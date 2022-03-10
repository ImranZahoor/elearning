package main

import (
	"elearning/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./quiz.db"), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to Open Database ")
	}

	g := gin.Default()
	g.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"intro": "e-learning api"})
	})
	module := g.Group("/module")
	{
		module.GET("/", controllers.GetAll)
		module.GET("/:id", controllers.Find)
	}

	subject := g.Group("/subject")
	{
		subject.GET("/", controllers.GetAll)
		subject.GET("/:id", controllers.Find)
	}

	g.Run("0.0.0.0:9090")
}
