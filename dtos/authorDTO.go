package dto

type AuthorDTO struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Books []string `json:"books"`
}
