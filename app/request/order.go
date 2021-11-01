package request

import (
	"errors"
	"github.com/spf13/cast"
	"github.com/zngue/go_helper/pkg/md5"
	"github.com/zngue/pay_provider/app/model"
	"github.com/zngue/pay_provider/app/service"
	"time"
)

//支付验签
type PaySign struct {
	Sign     string `form:"sign" binding:"required"`     //签名  appkey + appid + SignTime + SignStr  md5 加密 = Sign
	SignTime int    `form:"signTime" binding:"required"` //验签时间
	SignStr  string `form:"signStr" binding:"required"`  //签名字符串
}
type CreateOrderRequest struct {
	PaySign
	model.Order
}

func (m *CreateOrderRequest) SignCheck() error {
	if ( time.Now().Unix()-int64(m.PaySign.SignTime) )>60{
		return errors.New("签名验证超时")
	}
	var req  service.MerchantRequest
	req.AppKey=m.AppID
	detail, err2 := service.NewMerchant().Detail(&req)
	if err2 != nil {
		return err2
	}
	md5Str := md5.MD5(detail.MchAppKey + m.AppID + cast.ToString(m.SignTime) + m.SignStr)
	if md5Str != m.Sign {
		return errors.New("签名错误")
	}
	return nil
}

//6e684e2dd1fd541cc1a675aef8c0e21bzngd1fd541cc1a675ae123456signStr




