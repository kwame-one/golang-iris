package routes

import (
	"my-iris-app/controllers"
	"my-iris-app/convertors"
	"my-iris-app/repository"
	"my-iris-app/services"

	"github.com/kataras/iris/v12"
)

type BookRoute struct {
	controller controllers.BookController
	party      iris.Party
}

func InitBookRoute(p iris.Party) {
	route := &BookRoute{
		controller: *controllers.InitBookController(services.InitBookService(repository.InitBookRepository(), *convertors.InitBookConvertor())),
		party:      p,
	}

	route.party.Get("/", route.controller.GetAll)
	route.party.Post("/", route.controller.Post)
	route.party.Put("/{id:int}", route.controller.Put)
	route.party.Get("/{id:int}", route.controller.Get)
}
