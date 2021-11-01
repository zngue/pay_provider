package service

import (
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/zngue/pay_provider/app/model"
	"github.com/zngue/pay_provider/app/pay"
)

type Pay struct {
	Order *model.Order
	Merchant *model.Merchant
}
type PayInfo struct {
	Order *model.Order `json:"shop"`
	Jsapi *wechat.JSAPIPayParams `json:"payInfo"`
	CodeUrl string `json:"codeUrl"` 
}
type IPay interface {
	GetOrderPayInfo(req *OrderRequest) ( *PayInfo ,error)
}

func NewPay() IPay  {
	return &Pay{
		Order: model.NewOrder(),
		Merchant: model.NewMerchant(),
	}
}
func (p *Pay) GetParams(req *OrderRequest) ( err error){
	p.Order, err = NewOrder().Detail(req)
	if err != nil {
		return
	}
	var cmhReq MerchantRequest
	cmhReq.MchAppID=p.Order.AppID
	p.Merchant, err = NewMerchant().Detail(&cmhReq)
	if err != nil {
		return
	}
	return nil
}
func (p *Pay) GetOrderPayInfo(req *OrderRequest) ( *PayInfo ,error){
	if err := p.GetParams(req); err != nil {
		return nil, err
	}
	var returnData *PayInfo=&PayInfo{
		Order: p.Order,
	}
	if req.Status==1 {
		jsSDKOrder, errPay := new(pay.WechatPay).CreateJSSDKOrder(p.Merchant, p.Order)
		if errPay != nil {
			return nil, errPay
		}
		returnData.Jsapi=jsSDKOrder
		return returnData,nil
	}else{
		if codeUrl, err := new(pay.WechatPay).CreateNaviteOrder(p.Merchant, p.Order); err != nil {
			return nil, err
		}else{
			returnData.CodeUrl=codeUrl
			return returnData,nil
		}
	}
}







