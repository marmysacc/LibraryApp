package dto

type BookCreateDTO struct {
	Title    string `json:"title"`
	AuthorID string   `json:"author"`
	Genre    string `json:"genre"`
}
