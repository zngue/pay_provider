package order

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/pay_provider/app/request"
	"github.com/zngue/pay_provider/app/service"
	"math/rand"
	"strconv"
	"time"
)
func ApiBase() service.IOrder {
	return service.NewOrder()
}

func randomMath(end, start int64) int64 {
	randoms := rand.Int63n(end - start)
	return randoms + start
}
// Create  创建订单数据
func Create(ctx *gin.Context) {
	var req service.OrderRequest
	var data request.CreateOrderRequest
	if err := ctx.ShouldBind(&data); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	data.OrderNo="wx"+time.Now().Format("20060102150405")+
		strconv.Itoa(int(randomMath(999999,100000)))+
		strconv.Itoa(int(randomMath(999999,100000)))
	if err := data.SignCheck(); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	req.Data = &data.Order
	err := ApiBase().Add(&req)
	api.DataWithErr(ctx, err, data)
	return
}

// Edit  修改数据
func Edit(ctx *gin.Context) {
	var req service.OrderRequest
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
	req.Data = updateData
	err := ApiBase().Save(&req)
	api.DataWithErr(ctx, err, nil)
}

func List(ctx *gin.Context) {
	var req service.OrderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	list, err := ApiBase().List(&req)
	api.DataWithErr(ctx, err, list)
}
func Detail(ctx *gin.Context) {
	var req service.OrderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	list, err := ApiBase().Detail(&req)
	api.DataWithErr(ctx, err, list)
}
func Delete(ctx *gin.Context) {
	var req service.OrderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	err := ApiBase().Delete(&req)
	api.DataWithErr(ctx, err, nil)
}
func OrderPay(ctx *gin.Context)  {
	var req service.OrderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	info, err := service.NewPay().GetOrderPayInfo(&req)
	api.DataWithErr(ctx,err,info)
	return
}