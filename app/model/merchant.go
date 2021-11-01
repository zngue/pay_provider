package model

import (
	"gorm.io/gorm"
	"time"
)

type Merchant struct {
	ID int `gorm:"primaryKey;AUTO_INCREMENT;column:id;type:int(10);default:0" json:"id"`
	MchName string `gorm:"index;column:mch_name;type:varchar(50);default:" json:"mch_name" form:"mch_name"`
	MchAppID string `gorm:"uniqueIndex;column:mch_app_id;type:varchar(100);default:" json:"mch_app_id" form:"mch_app_id"`
	MchAppKey string `gorm:"column:mch_app_key;type:varchar(200);default:" json:"mch_app_key" form:"mch_app_key"`
	MchID string `gorm:"uniqueIndex;column:mch_id;type:varchar(20);default:" json:"mch_id" form:"mch_id"`
	SPAppID string `gorm:"index;column:sp_app_id;type:varchar(50);default:" json:"sp_app_id" form:"sp_app_id"`
	SPMchID string `gorm:"index;column:sp_mch_id;type:varchar(30);default:" json:"sp_mch_id" form:"sp_mch_id"`
	ApiKey2 string `gorm:"column:api_key2;type:varchar(60);default:" json:"api_key2" form:"api_key2"`
	ApiKey3 string `gorm:"column:api_key3;type:varchar(60);default:" json:"api_key3" form:"api_key3"`
	SerialNo string `gorm:"index;column:serial_no;type:varchar(100);default:" json:"serial_no" form:"serial_no"`
	ApiClientKeyPath string `gorm:"column:api_client_key_path;type:varchar(200);default:" json:"api_client_key_path" form:"api_client_key_path"`
	ApiClientCerPath string `gorm:"column:api_client_cer_path;type:varchar(200);default:" json:"api_client_cer_path" form:"api_client_cer_path"`
	SubAppID string `gorm:"index;column:sub_app_id;type:varchar(50);default:" json:"sub_app_id" form:"sub_app_id"`
	SubMchID string `gorm:"column:sub_mch_id;type:varchar(30);default:" json:"sub_mch_id" form:"sub_mch_id"`
	Email string `gorm:"index;column:email;type:varchar(50);default:" json:"email" form:"email"`
	IsServicePay int `gorm:"column:is_service_pay;type:tinyint(1);default:0" json:"is_service_pay" form:"is_service_pay"` // 0 普通支付 1 服务商支付  2 小程序
	NotifyUrl string `gorm:"column:notify_url;type:varchar(200);default:" json:"notify_url" form:"notify_url"`
	ApiClientCer12Path string `gorm:"column:api_client_cer12_path;type:varchar(200);default:" json:"api_client_cer12_path" form:"api_client_cer12_path"`
	Desc string `gorm:"column:desc;type:varchar(200);default:" json:"desc" form:"desc"`
	AddTime    int64 `gorm:"column:add_time;type:int(10);default:0" json:"addTime"`
	UpdateTime int64 `gorm:"column:update_time;type:int(10);default:0" json:"updateTime"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (i *Merchant) TableName() string {
	return "merchant"
}
func NewMerchant() *Merchant {
	return new(Merchant)
}
