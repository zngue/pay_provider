package service

import (
	"github.com/zngue/go_helper/pkg"
	"github.com/zngue/pay_provider/app/model"
)

type IPayLog interface {
	Add(request *PayLogRequest) error
	Save(request *PayLogRequest) error
	List(request *PayLogRequest) (*[]model.PayLog, error)
	Detail(request *PayLogRequest) (*model.PayLog, error)
	Delete(request *PayLogRequest) error
	GetModel() interface{}
}
type PayLog struct {
}
type PayLogRequest struct {
	pkg.CommonRequest
	ID int `form:"id" field:"id" where:"eq" default:"0"`
}

func NewPayLog() IPayLog {
	return new(PayLog)
}
func (p *PayLog) GetModel() interface{} {
	return model.NewPayLog()
}

// Add 添加
func (p *PayLog) Add(request *PayLogRequest) error {
	request.ReturnType = 3
	return pkg.MysqlConn.Model(p.GetModel()).Create(request.Data).Error
}

// Save 修改
func (p *PayLog) Save(request *PayLogRequest) error {
	request.ReturnType = 3
	return request.Init(pkg.MysqlConn.Model(p.GetModel()), *request).Updates(request.Data).Error
}
func (p *PayLog) List(request *PayLogRequest) (*[]model.PayLog, error) {
	var list []model.PayLog
	err := request.Init(pkg.MysqlConn.Model(p.GetModel()), *request).Find(&list).Error
	return &list, err
}

// Detail 详情
func (p *PayLog) Detail(request *PayLogRequest) (*model.PayLog, error) {
	var detail model.PayLog
	request.ReturnType = 3
	err := request.Init(pkg.MysqlConn.Model(p.GetModel()), *request).First(&detail).Error
	return &detail, err
}

// Delete 删除
func (p *PayLog) Delete(request *PayLogRequest) error {
	request.ReturnType = 3
	return request.Init(pkg.MysqlConn, *request).Delete(p.GetModel()).Error
}
