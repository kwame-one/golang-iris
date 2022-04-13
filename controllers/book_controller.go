package controllers

import (
	"my-iris-app/dtos"
	"my-iris-app/services"

	"github.com/kataras/iris/v12"
)

type BookController struct {
	service services.BookService
}

func InitBookController() *BookController {
	return &BookController{service: services.InitBookService()}
}

func (c *BookController) GetAll(ctx iris.Context) {

	books := c.service.GetAllBooks()

	ctx.JSON(books)
}

func (c *BookController) Post(ctx iris.Context) {
	var b dtos.BookRequest
	err := ctx.ReadJSON(&b)

	if err != nil {
		ctx.StopWithProblem(
			iris.StatusBadRequest,
			iris.NewProblem().Title("Book creation failure").DetailErr(err))

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
		ctx.StopWithProblem(
			iris.StatusBadRequest,
			iris.NewProblem().Title("Book creation failure").DetailErr(err))

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
