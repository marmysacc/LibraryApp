package service

import (
    "library-app/model"
    "library-app/repository"
)

type BookService interface {
    CreateBook(book *model.Book) error
    UpdateBook(book *model.Book) error
    DeleteBook(id uint) error
    GetBookByID(id uint) (*model.Book, error)
    GetAllBooks() ([]*model.Book, error)
}

type bookService struct {
    repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
    return &bookService{repo}
}

func (s *bookService) CreateBook(book *model.Book) error {
    return s.repo.Create(book)
}

func (s *bookService) UpdateBook(book *model.Book) error {
    return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
    return s.repo.Delete(id)
}

func (s *bookService) GetBookByID(id uint) (*model.Book, error) {
    return s.repo.GetByID(id)
}

func (s *bookService) GetAllBooks() ([]*model.Book, error) {
    return s.repo.GetAll()
}