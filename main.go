package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"spcplatform/dao"
	_ "spcplatform/docs"
	"spcplatform/middlewares"
)

// @title SPC Platform Swagger API
// @version 1.0
// @description SPC Platform Swagger API
// @contact.name caiwei
// @contact.email caiwei
// @BasePath /
// @securityDefinitions.apikey middlewares.ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//ac := middlewares.MakeAccessLog()
	ac := middlewares.MakeRotateAccessLog()
	defer ac.Close()

	var app = iris.New()
	app.UseRouter(recover.New(), ac.Handler)

	//init database
	dao.InitDataBase()
	//init api register
	registerApi(app)
	// init swagger docs
	registerSwaggerApi(app)
	app.Listen(":9999")
}
