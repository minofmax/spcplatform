package assets

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

type Assets interface {
	QueryByPage(ctx iris.Context)
	AddAssets(ctx iris.Context)
	DeleteAssets(ctx iris.Context)
}

type Host struct {
	IpAddress  string `json:"ipAddress"`
	IpType     string `json:"ipType"`
	Business   string `json:"business"`
	Owner      string `json:"owner"`
	Source     string `json:"source"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func NewHost(ctx iris.Context) Assets {
	return &Host{}
}

func (host *Host) AddAssets(ctx iris.Context) {
	fmt.Printf("route is %s, %s", ctx.FullRequestURI(), ctx.Path())
	ctx.JSON("123")
}

func (host *Host) QueryByPage(ctx iris.Context) {
	fmt.Println(ctx.GetCurrentRoute())
}

func (host *Host) DeleteAssets(ctx iris.Context) {
	fmt.Println(ctx.Host())
}

type Business struct {
	Name   string `json:"name"`
	Parent uint8  `json:"parent"`
	Child  uint8  `json:"child"`
}

func NewBusiness(ctx iris.Context) Assets {
	return &Business{}
}

func (business *Business) AddAssets(ctx iris.Context) {
	fmt.Printf("route is %s, %s", ctx.FullRequestURI(), ctx.Path())
	ctx.JSON("456")
}

func (business *Business) QueryByPage(ctx iris.Context) {
	fmt.Println(ctx.GetCurrentRoute())
}

func (business *Business) DeleteAssets(ctx iris.Context) {
	fmt.Println(ctx.Host())
}

type assetCreator func(ctx iris.Context) Assets

type Factory struct {
	creators map[string]assetCreator
}

func NewAssetFactory() *Factory {
	return &Factory{
		creators: map[string]assetCreator{
			"host":     NewHost,
			"business": NewBusiness,
		},
	}
}

func CreateAssets(ctx iris.Context) {
	factory := NewAssetFactory()
	var assetType = ctx.Params().Get("type")
	fmt.Printf("assets type %s", assetType)
	creator := factory.creators[assetType]
	assets := creator(ctx)
	assets.AddAssets(ctx)
}
