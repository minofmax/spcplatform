package views

import (
	"github.com/kataras/iris/v12"
	"spcplatform/services/assets"
)

func CreateAssets(ctx iris.Context) {
	factory := assets.NewAssetFactory()
	var assetType = ctx.Params().Get("type")
	creator := factory.Creators[assetType]
	assetsCreator := creator(ctx)
	assetsCreator.AddAssets(ctx)
}
