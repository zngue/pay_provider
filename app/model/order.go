package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID int `gorm:"primaryKey;column:id;type:int(10);default:0" json:"id"`
	AppID string `gorm:"index;column:app_id;type:varchar(100);default:" json:"app_id" form:"app_id"` //商户支付appid
	MchIDNo string `gorm:"index;column:mch_id_no;type:varchar(100);default:" json:"mch_id_no" form:"mch_id_no"`
	OrderNo string `gorm:"index;column:order_no;type:varchar(100);default:" json:"order_no" form:"order_no"`
	PaymentStatus int8 `gorm:"index;column:payment_status;type:tinyint(1);default:0" json:"payment_status" form:"payment_status"` //1 微信支付  2  支付宝支付
	TradeType string `gorm:"column:trade_type;type:varchar(50);default:" json:"trade_type" form:"trade_type"`
	Money int64 `gorm:"column:money;type:int(10);default:0" json:"money" form:"money"`
	SPOpenID string `gorm:"index;column:sp_open_id;type:varchar(80);default:" json:"sp_open_id" form:"sp_open_id"`
	SubOpenID string `gorm:"index;column:sub_open_id;type:varchar(80);default:" json:"sub_open_id" form:"sub_open_id"`
	PayStatus string `gorm:"column:pay_status;type:tinyint(1);default:0" json:"pay_status" form:"pay_status"`
	OrderTime int `gorm:"column:order_time;type:int(11);default:0" json:"order_time" form:"order_time"`
	NotifyUrl string `gorm:"column:notify_url;type:varchar(200);default:" json:"notify_url" form:"notify_url"`
	ReturnUrl string `gorm:"column:return_url;type:varchar(200);default:" json:"return_url" form:"return_url"`
	OrderDesc string `gorm:"column:order_desc;type:varchar(200);default:" json:"order_desc" form:"order_desc"`
	TransactionID string `gorm:"column:transaction_id;type:varchar(100);default:" json:"transaction_id" form:"transaction_id"`
	OrderPayTime int `gorm:"column:order_pay_time;type:int(11);default:0" json:"order_pay_time" form:"order_pay_time"`
	AddTime    int64 `gorm:"column:add_time;type:int(10);default:0" json:"addTime"`
	UpdateTime int64 `gorm:"column:update_time;type:int(10);default:0" json:"updateTime"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (o *Order) TableName() string {
	return "order"
}
func NewOrder() *Order {
	return new(Order)
}

type Resource struct {
	OriginalType   string `json:"original_type"`
	Algorithm      string `json:"algorithm"`
	Ciphertext     string `json:"ciphertext"`
	AssociatedData string `json:"associated_data"`
	Nonce          string `json:"nonce"`
}

type Callback struct {
	Id           string    `json:"id"`
	CreateTime   time.Time `json:"create_time"`
	ResourceType string    `json:"resource_type"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	Resource     *Resource `json:"resource"`
}
type Response struct {
	Mchid          string    `json:"mchid"`
	Appid          string    `json:"appid"`
	OutTradeNo     string    `json:"out_trade_no"`
	TransactionId  string    `json:"transaction_id"`
	TradeType      string    `json:"trade_type"`
	TradeState     string    `json:"trade_state"`
	TradeStateDesc string    `json:"trade_state_desc"`
	BankType       string    `json:"bank_type"`
	Attach         string    `json:"attach"`
	SuccessTime    time.Time `json:"success_time"`
	Payer          struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	Amount struct {
		Total         int    `json:"total"`
		PayerTotal    int    `json:"payer_total"`
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
}
