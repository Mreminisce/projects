package model

//  标签 帖子中间数据表
type PostTag struct {
	BaseModel
	PostId uint `gorm:"default:null comment '文章id'"` // id
	TagId  uint `gorm:"default:null comment '标签id'"` // 标签id
}
