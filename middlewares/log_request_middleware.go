package middlewares

import (
	"log"

	"github.com/kataras/iris/v12"
)

func Log(ctx iris.Context) {

	log.Println("method:", ctx.Method(), "path:", ctx.Path())

	ctx.Next()

}
