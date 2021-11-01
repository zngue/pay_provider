package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	wechat3 "github.com/go-pay/gopay/wechat/v3"
	"github.com/zngue/pay_provider/app/httplib"
	"github.com/zngue/pay_provider/app/model"
	"github.com/zngue/pay_provider/app/pay"
	"golang.org/x/sync/errgroup"
	"time"
)

type Notify struct {
	log      IPayLog
	merchant IMerchant
	pay      *pay.WechatPay
	order    IOrder
	ctx      *gin.Context
	clientV3 *wechat3.ClientV3
	merchantData *model.Merchant
}

type INotify interface {
	Init() error
	NotifyLog(str string,status int)
}

func NewNotify(ctx *gin.Context,merchant *model.Merchant) INotify {
	return &Notify{
		log:      NewPayLog(),
		pay:      pay.NewWechatPay(),
		order:    NewOrder(),
		ctx:      ctx,
		merchantData: merchant,
	}
}

func (n *Notify) GetWxKey (SerialNo string )  (*wechat3.PlatformCertItem ,*wechat3.SignInfo,error) {
	platformCerts, err := n.clientV3.GetPlatformCerts()
	if err != nil {
		return nil, nil, err
	}
	if len(platformCerts.Certs)>0 {
		for _, certItem := range platformCerts.Certs {
			if certItem!=nil {
				return  certItem ,platformCerts.SignInfo,nil
			}
		}
	}
	return nil,nil,errors.New("no matching SerialNo")
}

func  (n *Notify) V3ClientInit(merchant *model.Merchant)  error {
	client, err := new(pay.WechatPay).WechatService(merchant).WechatClient.V3Client()
	if err != nil {
		return err
	}
	n.clientV3=client
	return nil
}
func (n *Notify) Init() error {
	notify, noErr := wechat3.V3ParseNotify(n.ctx.Request)
	if noErr != nil {
		return noErr
	}
	if errs := n.V3ClientInit(n.merchantData); errs != nil {
		return errs
	}
	certItem, signInfo,wxErr := n.GetWxKey(n.merchantData.SerialNo)
	if wxErr != nil {
		return wxErr
	}
	notify.SignInfo=signInfo
	if err := notify.VerifySign(certItem.PublicKey); err != nil {
		return err
	}
	var (
		errs error
		notifyData *NotifyReturn
	)
	if n.merchantData.IsServicePay==1 {
		notifyData, errs = n.Service(notify)
	}else{
		notifyData, errs = n.Ordinary(notify)
	}
	if errs != nil {
		return errs
	}
	var wg errgroup.Group
	wg.Go(func() error {
		defer n.GoRouteRecoverService()
		return n.UpdateOrder(notifyData.OrderNo, notifyData.MchIDNo, notifyData.UpdateData)
	})
	wg.Go(func() error {
		defer n.GoRouteRecoverService()
		 n.httpNotify(notifyData.NotifyData)
		 return nil
	})
	if err := wg.Wait(); err != nil {
		go func() {
			defer n.GoRouteRecoverService()
			go n.NotifyLog(err.Error(),4)
		}()
		return err
	}
	return nil
}
type NotifyResponse struct {
	Code    int    `json:"code"` //200表示成功
	Message string `json:"message"`
}
func (n *Notify) httpNotify(data map[string]interface{}) error {
	jsonBody, er := httplib.Post(n.merchantData.NotifyUrl+"/"+n.merchantData.MchAppID).JSONBody(data)
	if er != nil {
		n.NotifyLog(er.Error(),3)
		return er
	}
	bytes, err := jsonBody.Bytes()
	if err != nil {
		return err
	}
	var notifyRsp NotifyResponse
	payErrs := json.Unmarshal(bytes, &notifyRsp)
	if payErrs != nil {
		return payErrs
	}
	sprintf := fmt.Sprintf("response:%s,request:%s", string(bytes), string(jsonBody.GetRequestJSONBody()))
	n.NotifyLog(sprintf,2)
	if notifyRsp.Code==200{
		return nil
	}else{
		return fmt.Errorf("回调失败")
	}
}
type NotifyReturn struct {
	UpdateData  map[string]interface{}
	NotifyData map[string]interface{}
	Openid string
	MchIDNo string
	OrderNo string
}

// Service 服务商支付回调
func (n *Notify) Service(req *wechat3.V3NotifyReq) (notifyData  *NotifyReturn,err error) {
	data, errZ := req.DecryptPartnerCipherText(n.merchantData.ApiKey3)
	if errZ != nil {
		return nil, errZ
	}
	var dataInfo NotifyReturn=NotifyReturn{
		UpdateData: map[string]interface{}{
			"pay_status":1,
			"sp_open_id":data.Payer.SpOpenid,
			"transaction_id":data.TransactionId,
			"trade_type":data.TradeType,
			"order_pay_time":time.Now().Unix(),
		},
		NotifyData: map[string]interface{}{
			"mch_id_no":data.Attach,
			"money":data.Amount.Total,
			"openid":data.Payer.SpOpenid,
			"trade_state":data.TradeState,
			"trade_state_desc":data.TradeStateDesc,
			"success_time":data.SuccessTime,
		},
		Openid: data.Payer.SpOpenid,
		OrderNo: data.OutTradeNo,
		MchIDNo: data.Attach,
	}
	go func() {
		defer n.GoRouteRecoverService()
		marshal, _ := json.Marshal(&data)
		n.NotifyLog(string(marshal),6)
	}()
	return &dataInfo,nil

}

// Ordinary 普通商户支付回调
func (n *Notify) Ordinary(req *wechat3.V3NotifyReq) ( notifyData  *NotifyReturn,err error) {
	data, errZ := req.DecryptCipherText(n.merchantData.ApiKey3)
	if errZ != nil {
		return nil, errZ
	}
	var dataInfo NotifyReturn=NotifyReturn{
		UpdateData: map[string]interface{}{
			"pay_status":1,
			"sp_open_id":data.Payer.Openid,
			"transaction_id":data.TransactionId,
			"trade_type":data.TradeType,
			"order_pay_time":time.Now().Unix(),
		},
		NotifyData: map[string]interface{}{
			"mch_id_no":data.Attach,
			"money":data.Amount.Total,
			"openid":data.Payer.Openid,
			"trade_state":data.TradeState,
			"trade_state_desc":data.TradeStateDesc,
			"success_time":data.SuccessTime,
		},
		Openid: data.Payer.Openid,
		OrderNo: data.OutTradeNo,
		MchIDNo: data.Attach,
	}
	go func() {
		defer n.GoRouteRecoverService()
		marshal, _ := json.Marshal(&data)
		n.NotifyLog(string(marshal),6)
	}()
	return &dataInfo,nil

}
func (n *Notify) UpdateOrder(OrderNo,MchIDNo string,data map[string]interface{}) error {
	var orderReq OrderRequest
	orderReq.OrderNo=OrderNo
	orderReq.MchIDNo=MchIDNo
	orderReq.Data=data
	return NewOrder().Save(&orderReq)
}
func (n *Notify) NotifyLog(str string,status int) {
	var logAdd PayLogRequest
	logAdd.Data=&model.PayLog{
		Content: str,
		Type: status,
	}
	n.log.Add(&logAdd)
}
func (n *Notify) GoRouteRecoverService() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}