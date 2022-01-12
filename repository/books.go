package repository

import (
	"deployment/models"

	"gorm.io/gorm"
)

type RepositoryBook interface {
	GetBooks() ([]models.Book, error)
	GetBookById(id int) (models.Book, error)
	CreateBook(book models.Book) (models.Book, error)
	DeleteBook(id int) error
	Updatebook(book models.Book) (models.Book, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepositoryBook(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetBooks() ([]models.Book, error) {
	var books []models.Book

	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *repository) GetBookById(id int) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) CreateBook(book models.Book) (models.Book, error) {

	err := r.db.Save(&book).Error
	if err != nil {
		return book, err
	}
	return book, err
}

func (r *repository) DeleteBook(id int) error {
	var book models.Book
	err := r.db.Delete(&book, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Updatebook(book models.Book) (models.Book, error) {

	err := r.db.Save(&book).Error
	if err != nil {
		return book, err
	}

	return book, err
}
