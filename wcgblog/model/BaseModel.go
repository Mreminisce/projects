package model

import "time"

// 使用BaseModel 代替grom.Model
type BaseModel struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"column:created_at;default:null comment '创建时间';"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:null comment '更新时间';"`
}
