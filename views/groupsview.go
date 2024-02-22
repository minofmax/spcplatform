package views

import (
	"github.com/kataras/iris/v12"
	"spcplatform/services/groups"
)

func QueryGroupPermissionsInPage(ctx iris.Context) {
	groups.QueryGroupPermissions(ctx)
}
