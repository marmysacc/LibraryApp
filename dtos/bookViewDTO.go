package dto

import "time"

// BookViewDTO represents a book with limited fields for the DTO.
type BookViewDTO struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	PublishedAt time.Time `json:"published_at"`
	AuthorName  string    `json:"author"`
}
