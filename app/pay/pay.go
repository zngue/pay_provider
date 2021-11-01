package pay

import (
	"errors"
	"github.com/go-pay/gopay"
	wechat2 "github.com/go-pay/gopay/wechat/v3"
	"github.com/zngue/go_helper/pkg/wechat"
	"github.com/zngue/pay_provider/app/model"
)
type WechatPay struct {
	WechatClient wechat.WechatClient
	Err error
}


func NewWechatPay() *WechatPay {
	return  new(WechatPay)
}
func  (w *WechatPay)  WechatService(mch *model.Merchant) *WechatPay  {
	wechatService:=&wechat.WechatService{
		WehcatConfig:&wechat.WehcatConfig{
			AppId:              mch.SPAppID,
			MchId: mch.SPMchID,
			Appkey:             mch.ApiKey2,
			ApiKey3: mch.ApiKey3,
			SerialNo: mch.SerialNo,
			ApiclientKeyPath: mch.ApiClientKeyPath,
			ApiclientCerPath: mch.ApiClientCerPath,
			ApiclientCer12Path: mch.ApiClientCer12Path,
			NotifyUrl: mch.NotifyUrl,
			IsProd: true,
		},
	}
	w.WechatClient=wechatService
	return w
}

func (w *WechatPay) CreateJSSDKOrder(mch *model.Merchant,order *model.Order) (*wechat2.JSAPIPayParams,error) {
	client, errs := w.WechatService(mch).WechatClient.V3Client()
	if errs != nil {
		return nil, errs
	}
	var (
		jsapi *wechat2.PrepayRsp
		err error
	)
	if mch.IsServicePay==1 {
		jsapi, err = client.V3PartnerTransactionJsapi(map[string]interface{}{
			"sp_appid":     mch.SPAppID,
			"sp_mchid":     mch.SPMchID,
			"sub_appid":    mch.SubAppID,
			"sub_mchid":    mch.SubMchID,
			"description":  order.OrderDesc,
			"out_trade_no": order.OrderNo,
			"attach":       order.MchIDNo,
			"notify_url":   mch.NotifyUrl,
			"amount": map[string]interface{}{
				"total": order.Money,
			},
			"payer": map[string]interface{}{
				"sp_openid":  order.SPOpenID,
				"sub_openid": order.SubOpenID,
			},
		})
	}else{
		jsapi, err = client.V3TransactionJsapi(map[string]interface{}{
			"appid":        mch.SPAppID,
			"mchid":        mch.SPMchID,
			"description":  order.OrderDesc,
			"out_trade_no": order.OrderNo,
			"attach":       order.MchIDNo,
			"notify_url":   mch.NotifyUrl,
			"amount": map[string]interface{}{
				"total": order.Money,
			},
			"payer": map[string]interface{}{
				"openid": order.SPOpenID,
			},
		})
	}
	if err != nil {
		return nil, err
	}
	if jsapi.Error!="" {
		return  nil,errors.New(jsapi.Error)
	}
	return client.PaySignOfJSAPI(mch.SPAppID, jsapi.Response.PrepayId)
}
func (w *WechatPay) CreateNaviteOrder(mch *model.Merchant,order *model.Order) (string,error) {

	client, errs := w.WechatService(mch).WechatClient.V3Client()
	if errs != nil {
		return "", errs
	}
	var (
		native *wechat2.NativeRsp
		err error
	)
	if mch.IsServicePay==1 {
		native, err = client.V3PartnerTransactionNative(map[string]interface{}{
			"sp_appid":     mch.SPAppID,
			"sp_mchid":     mch.SPMchID,
			"sub_appid":    mch.SubAppID,
			"sub_mchid":    mch.SubMchID,
			"description":  order.OrderDesc,
			"out_trade_no": order.OrderNo,
			"attach":       order.MchIDNo,
			"notify_url":   mch.NotifyUrl,
			"amount": map[string]interface{}{
				"total": order.Money,
			},
		})
	}else{
		native, err = client.V3TransactionNative(gopay.BodyMap{
			"appid":        mch.SPAppID,
			"mchid":        mch.SPMchID,
			"description":  order.OrderDesc,
			"out_trade_no": order.OrderNo,
			"attach":       order.MchIDNo,
			"notify_url":   mch.NotifyUrl,
			"amount": map[string]interface{}{
				"total": order.Money,
			},
		})
	}
	if err != nil {
		return "", err
	}
	if  native.Response!=nil &&  native.Response.CodeUrl!="" {
		return native.Response.CodeUrl, nil
	}
	return "", errors.New(native.Error)

}







