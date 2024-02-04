package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"spcplatform/dao"
	"spcplatform/middlewares"
)

func main() {
	ac := middlewares.MakeAccessLog()
	defer ac.Close()

	var app = iris.New()
	app.UseRouter(recover.New(), ac.Handler)

	//init database
	dao.InitDataBase()
	//init api register
	registerApi(app)

	app.Listen(":9999")
}
