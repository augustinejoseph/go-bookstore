package books

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Title       string    `json:"title" gorm:"not null"`
	Author      string    `json:"author" gorm:"not null"`
	PublishedAt time.Time `json:"published_at"`
	ISBN        string    `json:"isbn" gorm:"unique;not null"`
}
