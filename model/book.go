package model

import (
	"time"
)

type Book struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	AuthorID    string
	Author      Author    `json:"author"`
	Genre       string    `json:"genre"`
	PublishedAt time.Time `json:"published_at"`
}
