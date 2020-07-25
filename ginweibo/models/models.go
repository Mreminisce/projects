package models

import "time"

const (
	TrueTinyint  = 1
	FalseTinyint = 0
)

type BaseModel struct {
	ID        uint       `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index"` // *time.Time 支持 gorm 软删除
}
