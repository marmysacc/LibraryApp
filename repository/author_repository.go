package repository

import (
	"library-app/model"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	Create(author *model.Author) error
	Update(author *model.Author) error
	Delete(id uint) error
	GetByID(id uint) (*model.Author, error)
	GetAll() ([]*model.Author, error)
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) Create(author *model.Author) error {
	return r.db.Create(author).Error
}

func (r *authorRepository) Update(author *model.Author) error {
	return r.db.Save(author).Error
}

func (r *authorRepository) Delete(id uint) error {
	return r.db.Delete(&model.Author{}, id).Error
}

func (r *authorRepository) GetByID(id uint) (*model.Author, error) {
	var author model.Author
	if err := r.db.Preload("Books").First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *authorRepository) GetAll() ([]*model.Author, error) {
	var authors []*model.Author
	if err := r.db.Preload("Books").Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}
