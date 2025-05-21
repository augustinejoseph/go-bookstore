package main

import (
	"github.com/augustinejoseph/go-bookstore/internal/apps/books"
	"github.com/augustinejoseph/go-bookstore/internal/config"
	"github.com/augustinejoseph/go-bookstore/migrations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()
	// After ConnectDatabase() call
	if err := migrations.Migrate(config.DB); err != nil {
		panic("Failed to migrate database!")
	}

	router := gin.Default()

	router.Use(cors.Default())

	books.RegisterRoutes(router, config.DB)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.Run(":8080")
}
