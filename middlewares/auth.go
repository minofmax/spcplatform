package middlewares

import (
	"github.com/kataras/iris/v12"
	"spcplatform/services/models"
)

func CustomAuth() iris.Handler {
	// 拦截器实现cookie认证，用于自定义认证方法
	return func(ctx iris.Context) {
		ctx.Values().Set("framework", "iris")
		// before request, block request
		if ctx.Path() != "/api/books" {
			ctx.StatusCode(401)
			ctx.JSON(models.ResponseBody{Status: models.NotAuthorization, Msg: models.NotAuthorizationMsg})
			return
		}
		ctx.Next()
	}
}

func ApiKeyAuth() iris.Handler {
	return func(ctx iris.Context) {
		//fmt.Println()
		authorization := ctx.GetHeader("Authorization")
		if authorization != "123" {
			ctx.StatusCode(401)
			ctx.JSON(models.ResponseBody{Status: models.NotAuthorization, Msg: models.NotAuthorizationMsg})
			return
		}
		ctx.Next()
	}
}
