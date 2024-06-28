package service

import (
	dto "library-app/dtos"
	"library-app/model"
	"library-app/repository"
	"time"

	"github.com/google/uuid"
)

type AuthorService interface {
	CreateAuthor(authorDTO *dto.AuthorCreateDTO) error
	UpdateAuthor(author *model.Author) error
	DeleteAuthor(id uint) error
	GetAuthorByID(id uint) (*dto.AuthorViewDTO, error)
	GetAllAuthors() ([]*dto.AuthorViewDTO, error)
}

type authorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) AuthorService {
	return &authorService{repo}
}

func (s *authorService) CreateAuthor(authorDTO *dto.AuthorCreateDTO) error {
	author := s.mapToAuthorDTO(authorDTO)
	author.ID = uuid.New().String()
	author.CreatedAt = time.Now()
	return s.repo.Create(author)
}

func (s *authorService) UpdateAuthor(author *model.Author) error {
	return s.repo.Update(author)
}

func (s *authorService) DeleteAuthor(id uint) error {
	return s.repo.Delete(id)
}

func (s *authorService) GetAuthorByID(id uint) (*dto.AuthorViewDTO, error) {
	author, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.mapToAuthorViewDTO(author), nil
}

func (s *authorService) GetAllAuthors() ([]*dto.AuthorViewDTO, error) {
	authors, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var authorDTOs []*dto.AuthorViewDTO
	for _, author := range authors {
		authorDTOs = append(authorDTOs, s.mapToAuthorViewDTO(author))
	}
	return authorDTOs, nil
}

func (s *authorService) mapToAuthorViewDTO(author *model.Author) *dto.AuthorViewDTO {
	var bookTitles []string
	for _, book := range author.Books {
		bookTitles = append(bookTitles, book.Title)
	}
	return &dto.AuthorViewDTO{
		ID:    author.ID,
		Name:  author.Name,
		Books: bookTitles,
	}
}

func (s *authorService) mapToAuthorDTO(authorDTO *dto.AuthorCreateDTO) *model.Author {
	return &model.Author{
		Name: authorDTO.Name,
	}
}
