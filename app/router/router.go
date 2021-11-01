package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/pay_provider/app/api/merchant"
	"github.com/zngue/pay_provider/app/api/notify"
	"github.com/zngue/pay_provider/app/api/order"
)

func Router(api *gin.RouterGroup) {
	merchantRouters(api)
	orderRouters(api)
}
func merchantRouters(group *gin.RouterGroup) {

	merchantRouter := group.Group("merchant")
	{
		merchantRouter.POST("create", merchant.Create)
		merchantRouter.POST("edit", merchant.Edit)
		merchantRouter.POST("delete", merchant.Delete)
		merchantRouter.GET("list", merchant.List)
		merchantRouter.GET("detail", merchant.Detail)
	}
}
func orderRouters(group *gin.RouterGroup) {
	orderRouter := group.Group("order")
	{
		orderRouter.POST("create", order.Create)
		orderRouter.POST("edit", order.Edit)
		orderRouter.POST("delete", order.Delete)
		orderRouter.GET("list", order.List)
		orderRouter.GET("detail", order.Detail)
		orderRouter.GET("pay", order.OrderPay)
		orderRouter.POST("/AsyNotice/:appid",notify.AsyNotice)
	}
}