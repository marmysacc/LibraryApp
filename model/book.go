package model

import (
	"time"
)

type Book struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	AuthorID    uint
	Author      Author    `json:"author"`
	Genre       string    `json:"genre"`
	PublishedAt time.Time `json:"published_at"`
}
