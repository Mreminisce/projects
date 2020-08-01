package model

import "github.com/jinzhu/gorm"

// 用户订阅数据表
type Link struct {
	gorm.Model
	Name string `gorm:"type:varchar(10) not null comment '名称'"` //名称
	Url  string `gorm:"default: null comment '地址'"`             //地址
	Sort int    `gorm:"default:'0' comment '排序'"`               //排序
	View int    `gorm:"default:0 comment '访问次数'"`               //访问次数
}
