package service

import (
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/pay_provider/app/model"
)

type IMerchant interface {
	Add(request *MerchantRequest) error
	Save(request *MerchantRequest) error
	List(request *MerchantRequest) (*[]model.Merchant, error)
	Detail(request *MerchantRequest) (*model.Merchant, error)
	Delete(request *MerchantRequest) error
	GetModel() interface{}
}
type Merchant struct {

}

type MerchantRequest struct {
	pkg.CommonRequest
	ID int `form:"id" field:"id" where:"eq" default:"0"`
	AppKey string `form:"app_key" field:"app_key" where:"eq" default:""`
	SPAppID string `form:"sp_app_id" field:"sp_app_id" where:"eq" default:""`
	MchAppID string `form:"mch_app_id" field:"mch_app_id" where:"eq" default:""`
	
}

func NewMerchant() IMerchant {
	return new(Merchant)
}
func (m *Merchant) GetModel() interface{} {
	return model.NewMerchant()
}

// Add 添加
func (m *Merchant) Add(request *MerchantRequest) error {
	request.ReturnType = 3
	return pkg.MysqlConn.Model(m.GetModel()).Create(request.Data).Error
}

// Save 修改
func (m *Merchant) Save(request *MerchantRequest) error {
	request.ReturnType = 3
	return request.Init(pkg.MysqlConn.Model(m.GetModel()), *request).Updates(request.Data).Error
}
func (m *Merchant) List(request *MerchantRequest) (*[]model.Merchant, error) {
	var list []model.Merchant
	err := request.Init(pkg.MysqlConn.Model(m.GetModel()), *request).Find(&list).Error
	return &list, err
}

// Detail 详情
func (m *Merchant) Detail(request *MerchantRequest) (*model.Merchant, error) {
	var detail model.Merchant
	request.ReturnType = 3
	err := request.Init(pkg.MysqlConn.Model(m.GetModel()), *request).First(&detail).Error
	return &detail, err
}

// Delete 删除
func (m *Merchant) Delete(request *MerchantRequest) error {
	request.ReturnType = 3
	return request.Init(pkg.MysqlConn, *request).Delete(m.GetModel()).Error
}
