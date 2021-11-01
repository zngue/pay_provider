package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/pay_provider/app/model"
	"github.com/zngue/pay_provider/app/service"
	"time"
)

func ApiBase() service.IMerchant {
	return service.NewMerchant()
}

// Create  添加数据
func Create(ctx *gin.Context) {
	var req service.MerchantRequest
	var data model.Merchant
	time:=time.Now().Unix()
	data.AddTime=time
	data.UpdateTime=time
	if err := ctx.ShouldBind(&data); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	req.Data = &data
	err := ApiBase().Add(&req)
	api.DataWithErr(ctx, err, data)
	return
}

// Edit  修改数据
func Edit(ctx *gin.Context) {
	var req service.MerchantRequest
	formMap := ctx.PostFormMap("data")
	updateData := make(map[string]interface{})
	for key, val := range formMap {
		if key != "id" {
			updateData[key] = val
		}
	}
	//判断where 条件id是否存在  自行更改更新条件
	if id, ok := formMap["id"]; ok {
		newID := cast.ToInt(id)
		if newID <= 0 {
			api.Error(ctx, api.Msg("id 不能为空"))
			return
		}
		req.ID = newID
	}
	time:=time.Now().Unix()
	updateData["update_time"]=time
	req.Data = updateData
	err := ApiBase().Save(&req)
	api.DataWithErr(ctx, err, nil)
}

func List(ctx *gin.Context) {
	var req service.MerchantRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	list, err := ApiBase().List(&req)
	api.DataWithErr(ctx, err, list)
}
func Detail(ctx *gin.Context) {
	var req service.MerchantRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	list, err := ApiBase().Detail(&req)
	api.DataWithErr(ctx, err, list)
}
func Delete(ctx *gin.Context) {
	var req service.MerchantRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	err := ApiBase().Delete(&req)
	api.DataWithErr(ctx, err, nil)
}