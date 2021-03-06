/**
 * Copyright 2014 @ S1N1 Team.
 * name :
 * author : jarryliu
 * date : 2014-02-05 21:53
 * description :
 * history :
 */
package ols

import (
	"github.com/atnet/gof/web"
	"github.com/atnet/gof/web/mvc"
	"go2o/src/app/front/shop/ols/mos"
	"go2o/src/app/front/shop"
)

var (
	routes *mvc.Route = mvc.NewRoute(nil)
)

//处理请求
func Handle(ctx *web.Context) {
	switch GetBrownerDevice(ctx) {
		default:
		case shop.DevicePC:
		ctx.Items["device_view_dir"] = "ols"
		routes.Handle(ctx)
		case shop.DeviceTouchPad, shop.DeviceMobile:
		ctx.Items["device_view_dir"] = "mos"
		mos.Handle(ctx)
	}
}

//注册路由
func registerRoutes() {
	mc := &mainC{}

	sp := &ShoppingC{}
	pc := &PaymentC{}
	cc := &CartC{}
	uc := &UserC{}
	lc := &ListC{}

	routes.Register("buy", sp)
	routes.Register("shopping", sp)
	routes.Register("list", lc)
	routes.Register("cart", cc)
	routes.Register("user", uc)

	//处理错误
	routes.DeferFunc(func(ctx *web.Context) {
		if err, ok := recover().(error); ok {
			HandleCustomError(ctx.ResponseWriter, ctx, err)
		}
	})

	// 购物车接口
	routes.Add("/cart_api_v1", cc.CartApiHandle)
	// 支付
	routes.Add("^/pay/create", pc.Create)
	// 首页
	routes.Add("/", mc.Index)
	routes.Add("/user/jump_m", uc.JumpToMCenter)
}


func init(){
	registerRoutes()
}