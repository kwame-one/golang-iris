package convertors

import (
	"my-iris-app/dtos"
	"my-iris-app/models"
)

type BookConvertor struct{}

func InitBookConvertor() *BookConvertor {
	return &BookConvertor{}
}

func (convertor *BookConvertor) ToBook(bookDto dtos.BookDto) models.Book {
	return models.Book{
		Id:     bookDto.Id,
		Title:  bookDto.Title,
		Author: bookDto.Author,
	}
}

func (convertor *BookConvertor) ToBookDto(book models.Book) dtos.BookDto {
	return dtos.BookDto{
		Id:     book.Id,
		Title:  book.Title,
		Author: book.Author,
	}
}

func (convertor *BookConvertor) ParseBookRequest(bookRequest dtos.BookRequest) models.Book {
	return models.Book{
		Title:  bookRequest.Title,
		Author: bookRequest.Author,
	}
}
