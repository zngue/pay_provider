package notify

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/pay_provider/app/service"
)
func AsyNotice(ctx *gin.Context) {
	appid := ctx.Param("appid")
	var mch service.MerchantRequest
	mch.MchAppID=appid
	merchant, err := service.NewMerchant().Detail(&mch)
	if err != nil {
		api.WeChatPayError(ctx)
		return
	}
	notifyF := service.NewNotify(ctx, merchant)
	if err := notifyF.Init(); err != nil {
		notifyF.NotifyLog(err.Error(),5)
		api.WeChatPayError(ctx)
		return
	}
	notifyF.NotifyLog("支付成功",1)
	api.WeChatPaySuccess(ctx)
}


