package service

import (
	dto "library-app/dtos"
	"library-app/model"
	"library-app/repository"
)

type AuthorService interface {
	CreateAuthor(author *model.Author) error
	UpdateAuthor(author *model.Author) error
	DeleteAuthor(id uint) error
	GetAuthorByID(id uint) (*dto.AuthorDTO, error)
	GetAllAuthors() ([]*dto.AuthorDTO, error)
}

type authorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) AuthorService {
	return &authorService{repo}
}

func (s *authorService) CreateAuthor(author *model.Author) error {
	return s.repo.Create(author)
}

func (s *authorService) UpdateAuthor(author *model.Author) error {
	return s.repo.Update(author)
}

func (s *authorService) DeleteAuthor(id uint) error {
	return s.repo.Delete(id)
}

func (s *authorService) GetAuthorByID(id uint) (*dto.AuthorDTO, error) {
	author, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.mapToAuthorDTO(author), nil
}

func (s *authorService) GetAllAuthors() ([]*dto.AuthorDTO, error) {
	authors, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var authorDTOs []*dto.AuthorDTO
	for _, author := range authors {
		authorDTOs = append(authorDTOs, s.mapToAuthorDTO(author))
	}
	return authorDTOs, nil
}

func (s *authorService) mapToAuthorDTO(author *model.Author) *dto.AuthorDTO {
	var bookTitles []string
	for _, book := range author.Books {
		bookTitles = append(bookTitles, book.Title)
	}
	return &dto.AuthorDTO{
		ID:    author.ID,
		Name:  author.Name,
		Books: bookTitles,
	}
}
