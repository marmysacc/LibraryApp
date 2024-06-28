package service

import (
	dto "library-app/dtos"
	"library-app/model"
	"library-app/repository"
	"time"

	"github.com/google/uuid"
)

type BookService interface {
	CreateBook(bookDTO *dto.BookCreateDTO) error
	UpdateBook(book *model.Book) error
	DeleteBook(id string) error
	GetBookByID(id string) (*dto.BookViewDTO, error)
	GetAllBooks() ([]*dto.BookViewDTO, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo}
}

// func (s *bookService) CreateBook(book *dto.BookCreateDTO) error {
// 	return s.repo.Create(book)
// }

func (s *bookService) CreateBook(bookDTO *dto.BookCreateDTO) error {
	book := s.mapToBookDTO(bookDTO)
	book.ID = uuid.New().String()
	book.PublishedAt = time.Now()
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(book *model.Book) error {
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id string) error {
	return s.repo.Delete(id)
}

func (s *bookService) GetBookByID(id string) (*dto.BookViewDTO, error) {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.mapToBookViewDTO(book), nil
}

func (s *bookService) GetAllBooks() ([]*dto.BookViewDTO, error) {
	books, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var bookDTOs []*dto.BookViewDTO
	for _, book := range books {
		bookDTOs = append(bookDTOs, s.mapToBookViewDTO(book))
	}
	return bookDTOs, nil
}

func (s *bookService) mapToBookViewDTO(book *model.Book) *dto.BookViewDTO {
	return &dto.BookViewDTO{
		ID:          book.ID,
		Title:       book.Title,
		Genre:       book.Genre,
		PublishedAt: book.PublishedAt,
		AuthorName:  book.Author.Name,
	}
}

func (s *bookService) mapToBookDTO(bookDTO *dto.BookCreateDTO) *model.Book {
	return &model.Book{
		Title:    bookDTO.Title,
		Genre:    bookDTO.Genre,
		AuthorID: bookDTO.AuthorID,
	}
}
