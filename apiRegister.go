package main

import (
	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/kataras/iris/v12"
	"spcplatform/middlewares"
	"spcplatform/services/books"
	"spcplatform/views"
)

func registerSwaggerApi(app *iris.Application) {
	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("/swagger/swagger.json"),
		swagger.DeepLinking(true),
		swagger.Prefix("/swagger"),
		swagger.SetTheme(swagger.FeelingBlue),
	)
	app.Get("/swagger", swaggerUI)
	app.Get("/swagger/{any:path}", swaggerUI)
}

func registerApi(app *iris.Application) {
	apiParty := app.Party("/api")
	apiParty.Use(iris.Compression,
		//middlewares.CustomAuth(),
		middlewares.ApiKeyAuth(),
		middlewares.CustomRecover)
	registerBooksApi(apiParty)
	registerAssetsApi(apiParty)
	registerGroupsApi(apiParty)
}

func registerGroupsApi(partyAPI iris.Party) {
	partyAPI = partyAPI.Party("/groups/")
	{
		partyAPI.Get("/query", views.QueryGroupPermissionsInPage)
	}
}

func registerBooksApi(partyAPI iris.Party) {
	partyAPI = partyAPI.Party("/books/")
	{
		partyAPI.Get("/", books.List)
		partyAPI.Post("/", books.Create)
	}
}

func registerAssetsApi(partyAPI iris.Party) {
	partyAPI = partyAPI.Party("/assets/")
	{
		partyAPI.Get("/{type:path}", views.CreateAssets)
	}
}
