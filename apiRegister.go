package main

import (
	"github.com/kataras/iris/v12"
	"spcplatform/middlewares"
	"spcplatform/services/assets"
	"spcplatform/services/books"
)

func registerApi(app *iris.Application) {
	apiParty := app.Party("/api")
	apiParty.Use(iris.Compression,
		middlewares.CustomAuth(),
		middlewares.CustomRecover)
	registerBooksApi(apiParty)
	registerAssetsApi(apiParty)
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
		partyAPI.Get("/{type:path}", assets.CreateAssets)
	}
}
