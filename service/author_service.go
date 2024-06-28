package service

import (
	"library-app/model"
	"library-app/repository"
)

type AuthorService interface {
	CreateAuthor(author *model.Author) error
	UpdateAuthor(author *model.Author) error
	DeleteAuthor(id uint) error
	GetAuthorByID(id uint) (*model.Author, error)
	GetAllAuthors() ([]*model.Author, error)
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

func (s *authorService) GetAuthorByID(id uint) (*model.Author, error) {
	return s.repo.GetByID(id)
}

func (s *authorService) GetAllAuthors() ([]*model.Author, error) {
	return s.repo.GetAll()
}
