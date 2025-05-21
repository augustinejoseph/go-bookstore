package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB){
	bookController := NewBookController(db)
	books := r.Group("/books")
	{
		books.GET("/", bookController.ListBooks)
		books.GET("/:id", bookController.GetBook)
		books.POST("/", bookController.CreateBook)
	}
}