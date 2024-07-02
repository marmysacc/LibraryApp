package service

import (
	"fmt"
	dto "library-app/dtos"
	"library-app/model"
	"library-app/repository"
	"time"

	"github.com/google/uuid"
)

type BookService interface {
	CreateBook(bookDTO *dto.BookCreateDTO) error
	UpdateBook(book *dto.BookCreateDTO, id string) error
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

func (s *bookService) CreateBook(bookDTO *dto.BookCreateDTO) error {
	book := s.mapToBookDTO(bookDTO)
	book.ID = uuid.New().String()
	book.PublishedAt = time.Now()
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(bookDTO *dto.BookCreateDTO, id string) error {
	currentBook, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	updatedBook := s.mapToBookUpdateDTO(bookDTO, currentBook)

	return s.repo.Update(updatedBook)
}

func (s *bookService) DeleteBook(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid UUID format for id: %s", id)
	}

	exists, err := s.repo.ExistsByID(id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("book with id %s does not exist", id)
	}

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

func (s *bookService) mapToBookUpdateDTO(bookDTO *dto.BookCreateDTO, currentBook *model.Book) *model.Book {
	updatedBook := *currentBook

	if bookDTO.Title != "" {
		updatedBook.Title = bookDTO.Title
	}
	if bookDTO.AuthorID != "" {
		updatedBook.AuthorID = bookDTO.AuthorID
	}
	if bookDTO.Genre != "" {
		updatedBook.Genre = bookDTO.Genre
	}

	return &updatedBook
}
