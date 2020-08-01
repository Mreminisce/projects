package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// user 数据表
type User struct {
	gorm.Model
	Email       string    `gorm:"unique_index;default:null comment '邮箱'"` //邮箱
	Password    string    `gorm:"default:null comment '密码'"`              //密码
	VerifyState string    `gorm:"default:'0' comment '邮箱验证状态'"`           //邮箱验证状态
	SecretKey   string    `gorm:"default:null  comment '密钥'"`             //密钥
	OutTime     time.Time `gorm:"default:null  comment '过期时间'"`           //过期时间
	IsAdmin     bool      `gorm:"default:1 comment '是否是管理员'"`             //是否是管理员
	AvatarUrl   string    `gorm:"default:null comment '头像链接'"`            // 头像链接
	LockState   bool      `gorm:"default:'0' comment '锁定状态'"`             //锁定状态
}
