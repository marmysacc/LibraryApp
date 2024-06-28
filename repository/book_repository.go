package repository

import (
	"library-app/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *model.Book) error
	Update(book *model.Book) error
	Delete(id string) error
	GetByID(id string) (*model.Book, error)
	GetAll() ([]*model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) Create(book *model.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) Update(book *model.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id string) error {
	return r.db.Delete(&model.Book{}, "id = ?", id).Error
}


func (r *bookRepository) GetByID(id string) (*model.Book, error) {
	var book model.Book
	if err := r.db.Preload("Author").First(&book, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}


func (r *bookRepository) GetAll() ([]*model.Book, error) {
	var books []*model.Book
	if err := r.db.Preload("Author").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
