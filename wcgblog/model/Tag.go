package model

type Tag struct {
	BaseModel
	Name  string `gorm:"default:null comment '标签名称'"`
	Total int    `gorm:"-"`
}
