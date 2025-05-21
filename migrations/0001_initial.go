package migrations

import (
	"github.com/augustinejoseph/go-bookstore/internal/apps/books"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&books.Book{}); err != nil {
		return err
	}


	return nil
}
