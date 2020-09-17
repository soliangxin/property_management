package routers

import (
	"github.com/astaxie/beego"
	"server/controllers"
)

func init() {
	loginController := &controllers.LoginController{}
	beego.Router("/user/login", loginController, "*:Login")
	beego.Router("/user/info", loginController, "*:GetInfo")
	beego.Router("/user/logout", loginController, "*:Logout")

	// 小区管理
	community := &controllers.CommunityControllers{}
	beego.Router("/community/list", community, "*:GetCommunityList")

	// 房产管理
	house := &controllers.HouseControllers{}
	beego.Router("/house/list", house, "*:GetHouseList")

	// 住户管理
	tenement := &controllers.TenementControllers{}
	beego.Router("/tenement/list", tenement, "*:GetTenementList")

	// 停车位管理
	stall := &controllers.StallControllers{}
	beego.Router("/stall/list", stall, "*:GetStallList")

	// 收费管理
	charge := &controllers.ChargeControllers{}
	beego.Router("/change/itemList", charge, "*:GetChangeItemList")
	beego.Router("/change/recordList", charge, "*:GetChangeRecordList")

	// 值班管理
	watch := &controllers.WatchControllers{}
	beego.Router("/watch/list", watch, "*:GetWatchList")

	// 资产管理相关操作
	asset := &controllers.AssetControllers{}
	beego.Router("/asset/list", asset, "*:GetAssetList")

	// 请求
	beego.Router("/sockjs-node/*", &controllers.SockJsController{})
}
