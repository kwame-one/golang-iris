package services

import (
	"my-iris-app/convertors"
	"my-iris-app/dtos"
	"my-iris-app/repository"
)

func InitBookService() BookService {
	return &BookServiceImpl{
		repo:      repository.InitBookRepository(),
		convertor: *convertors.InitBookConvertor(),
	}
}

type BookServiceImpl struct {
	repo      repository.BookRepository
	convertor convertors.BookConvertor
}

// GetBook implements BookService
func (service *BookServiceImpl) GetBook(id int) (dtos.BookDto, error) {
	book, err := service.repo.GetBook(id)

	if err != nil {
		return dtos.BookDto{}, err
	}

	return service.convertor.ToBookDto(book), nil
}

// CreateBook implements BookService
func (service *BookServiceImpl) CreateBook(bookRequest dtos.BookRequest) (dtos.BookDto, error) {

	book := service.convertor.ParseBookRequest(bookRequest)

	newBook := service.repo.CreateBook(book)

	return service.convertor.ToBookDto(newBook), nil

}

// GetAllBooks implements BookService
func (service *BookServiceImpl) GetAllBooks() []dtos.BookDto {
	books := service.repo.GetAllBooks()

	bookDtos := make([]dtos.BookDto, len(books))

	for index, book := range books {
		bookDtos[index] = service.convertor.ToBookDto(book)
	}

	return bookDtos
}

// UpdateBook implements BookService
func (service *BookServiceImpl) UpdateBook(id int, bookRequest dtos.BookRequest) (dtos.BookDto, error) {
	book, err := service.repo.GetBook(id)

	if err != nil {
		return dtos.BookDto{}, err
	}

	book.Title = bookRequest.Title

	updatedBook := service.repo.UpdateBook(id, book)

	return service.convertor.ToBookDto(updatedBook), nil
}
