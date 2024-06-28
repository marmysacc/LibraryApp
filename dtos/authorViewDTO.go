package dto

type AuthorViewDTO struct {
	ID    string     `json:"id"`
	Name  string   `json:"name"`
	Books []string `json:"books"`
}
