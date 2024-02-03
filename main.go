package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	var app = iris.New()
	app.UseRouter(recover.New())
	registerApi(app)
	app.Listen(":9999")
}
