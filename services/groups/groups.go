package groups

import (
	"github.com/kataras/iris/v12"
	"spcplatform/dao"
)

func QueryGroupPermissions(ctx iris.Context) {
	pageSize := ctx.URLParamIntDefault("pageSize", 10)
	pageNum := ctx.URLParamIntDefault("pageNum", 1)
	offset := (pageNum - 1) * pageSize
	page := dao.QueryGroupPermissionDataInPage(pageSize, offset)
	ctx.JSON(page)
}
