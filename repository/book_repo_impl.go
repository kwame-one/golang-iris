package repository

import (
	"errors"
	"my-iris-app/models"
)

type BookRepositoryImpl struct {
	books []models.Book
}

func InitBookRepository() BookRepository {
	return &BookRepositoryImpl{
		books: make([]models.Book, 0),
	}
}

func (repo *BookRepositoryImpl) GetBook(id int) (models.Book, error) {
	book := models.Book{
		Id:    0,
		Title: "",
	}

	for _, b := range repo.books {
		if b.Id == id {
			book = b
			break
		}
	}

	if book.Id == 0 {
		return models.Book{}, errors.New("book not found")
	}

	return book, nil
}

func (repo *BookRepositoryImpl) GetAllBooks() []models.Book {
	return repo.books
}

func (repo *BookRepositoryImpl) CreateBook(book models.Book) models.Book {

	newBook := models.Book{
		Id:    len(repo.books) + 1,
		Title: book.Title,
	}

	repo.books = append(repo.books, newBook)

	return newBook
}

func (repo *BookRepositoryImpl) UpdateBook(id int, book models.Book) models.Book {

	currentBook := repo.books[id-1]

	currentBook.Title = book.Title

	repo.books[id-1] = currentBook

	return repo.books[id-1]
}
