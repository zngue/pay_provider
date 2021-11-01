package service

import (
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/pay_provider/app/model"
)

type IOrder interface {
	Add(request *OrderRequest) error
	Save(request *OrderRequest) error
	List(request *OrderRequest) (*[]model.Order, error)
	Detail(request *OrderRequest) (*model.Order, error)
	Delete(request *OrderRequest) error
	GetModel() interface{}
}
type Order struct {

}
type OrderRequest struct {
	pkg.CommonRequest
	ID int `form:"id" field:"id" where:"eq" default:"0"`
	MchIDNo string `form:"mch_id_no" field:"mch_id_no" where:"eq" default:""`
	OpenID string `form:"open_id" field:"open_id" where:"eq" default:""`
	PayStatus int `form:"pay_status" field:"pay_status" where:"eq" default:"-1"`
	OrderNo string `form:"order_no" field:"order_no" where:"eq" default:""`
	Status int `form:"status"` //1 jsapi  2 扫码支付

}

func NewOrder() IOrder {
	return new(Order)
}
func (o *Order) GetModel() interface{} {
	return model.NewOrder()
}

// Add 添加
func (o *Order) Add(request *OrderRequest) error {
	request.ReturnType = 3
	return pkg.MysqlConn.Model(o.GetModel()).Create(request.Data).Error
}

// Save 修改
func (o *Order) Save(request *OrderRequest) error {
	request.ReturnType = 3
	return request.Init(pkg.MysqlConn.Model(o.GetModel()), *request).Updates(request.Data).Error
}
func (o *Order) List(request *OrderRequest) (*[]model.Order, error) {
	var list []model.Order
	err := request.Init(pkg.MysqlConn.Model(o.GetModel()), *request).Find(&list).Error
	return &list, err
}

// Detail 详情
func (o *Order) Detail(request *OrderRequest) (*model.Order, error) {
	var detail model.Order
	request.ReturnType = 3
	err := request.Init(pkg.MysqlConn.Model(o.GetModel()), *request).First(&detail).Error
	return &detail, err
}

// Delete 删除
func (o *Order) Delete(request *OrderRequest) error {
	request.ReturnType = 3
	return request.Init(pkg.MysqlConn, *request).Delete(o.GetModel()).Error
}
