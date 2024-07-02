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
	UpdateAuthor(authorDTO *dto.AuthorCreateDTO, id string) error
	DeleteAuthor(id string) error
	GetAuthorByID(id string) (*dto.AuthorViewDTO, error)
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

func (s *authorService) UpdateAuthor(authorDTO *dto.AuthorCreateDTO, id string) error {
	currentAuthor, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	updatedAuthor := s.mapToAuthorUpdateDTO(authorDTO, currentAuthor)
	return s.repo.Update(updatedAuthor)
}

func (s *authorService) DeleteAuthor(id string) error {
	return s.repo.Delete(id)
}

func (s *authorService) GetAuthorByID(id string) (*dto.AuthorViewDTO, error) {
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

func (s *authorService) mapToAuthorUpdateDTO(authorDTO *dto.AuthorCreateDTO, currentAuthor *model.Author) *model.Author {
	updatedAuthor := *currentAuthor

	if authorDTO.Name != "" {
		updatedAuthor.Name = authorDTO.Name
	}
	
	return &updatedAuthor
}
