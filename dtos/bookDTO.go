package dto

import "time"

// BookDTO represents a book with limited fields for the DTO.
type BookDTO struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	PublishedAt time.Time `json:"published_at"`
	AuthorName  string    `json:"author"`
}
