package groups

import (
	"github.com/kataras/iris/v12"
	"spcplatform/dao"
)

func QueryGroupPermissions(ctx iris.Context) {
	pageSize, err := ctx.Params().GetInt("pageSize")
	if err != nil {
		pageSize = 10
	}
	pageNum, err := ctx.Params().GetInt("pageNum")
	if err != nil {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize
	page := dao.QueryGroupPermissionDataInPage(pageSize, offset)
	ctx.JSON(page)
}
