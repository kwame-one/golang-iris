package main

import (
	"my-iris-app/routes"

	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Validator = validator.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
	})

	app.UseRouter(crs)

	books := app.Party("/books")

	routes.InitBookRoute(books)

	app.Listen(":8080")
}
