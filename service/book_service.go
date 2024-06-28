package service

import (
	dto "library-app/dtos"
	"library-app/model"
	"library-app/repository"
)

type BookService interface {
	CreateBook(book *model.Book) error
	UpdateBook(book *model.Book) error
	DeleteBook(id uint) error
	GetBookByID(id uint) (*dto.BookDTO, error)
	GetAllBooks() ([]*dto.BookDTO, error)
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

func (s *bookService) GetBookByID(id uint) (*dto.BookDTO, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.mapToBookDTO(book), nil
}

func (s *bookService) GetAllBooks() ([]*dto.BookDTO, error) {
	books, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var bookDTOs []*dto.BookDTO
	for _, book := range books {
		bookDTOs = append(bookDTOs, s.mapToBookDTO(book))
	}
	return bookDTOs, nil
}

func (s *bookService) mapToBookDTO(book *model.Book) *dto.BookDTO {
	return &dto.BookDTO{
		ID:          book.ID,
		Title:       book.Title,
		Genre:       book.Genre,
		PublishedAt: book.PublishedAt,
		AuthorName:  book.Author.Name,
	}
}
