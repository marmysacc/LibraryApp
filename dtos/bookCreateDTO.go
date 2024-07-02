package dto

type BookCreateDTO struct {
	Title    string `json:"title" validate:"required,min=2,max=100"`
	AuthorID string `json:"author" validate:"required,uuid"`
	Genre    string `json:"genre" validate:"required,min=2,max=25"`
}
