package repository

import (
	"github.com/iyiola-dev/go-graphql/app/models"
	"github.com/iyiola-dev/go-graphql/graph/model"
	"gorm.io/gorm"
)

type BookRepository interface{}

type BookService struct {
	Db *gorm.DB
}

func NewBookService(db *gorm.DB) *BookService {
	return &BookService{
		Db: db,
	}
}

func (b *BookService) CreateBook(bookInput *model.BookInput) (*models.Book, error) {
	book := &models.Book{
		Title:     bookInput.Title,
		Author:    bookInput.Author,
		Publisher: bookInput.Publisher,
	}
	err := b.Db.Create(&book).Error

	return book, err
}
