package model

type Comment struct {
	BaseModel
	UserID    uint
	Content   string `gorm:"default:null comment '内容'"`
	PostID    uint   `gorm:"column:post_id; default:null comment '文章id'"`
	ReadState bool   `gorm:"column:read_state;default:'0' comment '阅读状态'"`
	AvatarUrl string `gorm:"-"`
}
