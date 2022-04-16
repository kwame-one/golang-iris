package controllers

import (
	"my-iris-app/dtos"
	"my-iris-app/services"
	customValidator "my-iris-app/validators"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type BookController struct {
	service services.BookService
}

func InitBookController(bookService services.BookService) *BookController {
	return &BookController{service: bookService}
}

func (c *BookController) GetAll(ctx iris.Context) {

	books := c.service.GetAllBooks()

	ctx.JSON(books)
}

func (c *BookController) Post(ctx iris.Context) {
	var b dtos.BookRequest
	err := ctx.ReadJSON(&b)

	if err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			errors := customValidator.CustomMessages(errs)

			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
				Title("Validator Error").
				Key("errors", errors))

			return

		}

		ctx.StopWithStatus(iris.StatusInternalServerError)
		return

	}

	book, _ := c.service.CreateBook(b)

	ctx.StatusCode(iris.StatusCreated)

	ctx.JSON(book)
}

func (c *BookController) Get(ctx iris.Context) {
	bookId, _ := ctx.Params().GetInt("id")

	book, err := c.service.GetBook(bookId)

	if err != nil {
		ctx.StopWithProblem(
			iris.StatusNotFound,
			iris.NewProblem().Title(err.Error()).DetailErr(err))

		return
	}

	ctx.JSON(book)
}

func (c *BookController) Put(ctx iris.Context) {
	bookId, _ := ctx.Params().GetInt("id")

	var b dtos.BookRequest
	err := ctx.ReadJSON(&b)

	if err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			errors := customValidator.CustomMessages(errs)

			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
				Title("Validator Error").
				Key("errors", errors))

			return

		}

		ctx.StopWithStatus(iris.StatusInternalServerError)
		return

	}

	book, err := c.service.UpdateBook(bookId, b)

	if err != nil {
		ctx.StopWithProblem(
			iris.StatusNotFound,
			iris.NewProblem().Title("Book update failure").DetailErr(err))

		return
	}

	ctx.JSON(book)
}
