package services

import "my-iris-app/dtos"

type BookService interface {
	GetAllBooks() []dtos.BookDto

	CreateBook(bookRequest dtos.BookRequest) (dtos.BookDto, error)

	UpdateBook(id int, bookRequest dtos.BookRequest) (dtos.BookDto, error)

	GetBook(id int) (dtos.BookDto, error)
}
