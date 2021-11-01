package model

import (
	"gorm.io/gorm"
	"time"
)

type PayLog struct { //支付日志
	ID        int            `gorm:"column:id;primaryKey;autoIncrement;type:int;size:11;" json:"id" form:"id"`
	Content   string         `gorm:"column:content;type:text not null;comment:;" json:"content" form:"content"`
	Type      int            `gorm:"column:type;type:int not null;size:2;" json:"type" form:"type" binding:"type"` //1 表示本地日志 2表示通知商户日志 3,本地执行错误日志
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at" form:"deleted_at"`
}
func (p *PayLog) TableName() string {
	return "pay_log"
}
func NewPayLog() *PayLog {
	return new(PayLog)
}
