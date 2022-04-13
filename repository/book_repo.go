package repository

import "my-iris-app/models"

type BookRepository interface {
	GetAllBooks() []models.Book

	CreateBook(book models.Book) models.Book

	UpdateBook(id int, book models.Book) models.Book

	GetBook(id int) (models.Book, error)
}
