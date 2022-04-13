package main

import (
	"my-iris-app/routes"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	books := app.Party("/books")

	routes.InitBookRoute(books)

	app.Listen(":8080")
}
