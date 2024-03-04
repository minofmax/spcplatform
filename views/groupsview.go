package views

import (
	"github.com/kataras/iris/v12"
	"spcplatform/services/groups"
)

// QueryGroupPermissionsInPage 分页查询permission信息
// @Tags Permission
// @Summary 查询permission信息
// @Accept json
// @Produce json
// @Param pageNum query integer no "分页页数"
// @Param pageSize query integer no "分页大小"
// @Success 200 "返回成功"
// @Failure 400 {object} models.ResponseBody "返回body"
// @Router /api/groups/query [get]
// @Security middlewares.ApiKeyAuth
func QueryGroupPermissionsInPage(ctx iris.Context) {
	groups.QueryGroupPermissions(ctx)
}
