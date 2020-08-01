package model

type Post struct {
	BaseModel
	Title        string     `gorm:"default:null comment '标题'"`
	Body         string     `gorm:"default:null comment '主体'"`
	View         int        `gorm:"default:0 comment '浏览次数'"`
	IsPublished  bool       `gorm:"default:'0' comment '是否发表'"`
	Tags         []*Tag     `gorm:"-"`
	Comments     []*Comment `gorm:"-"`
	CommentTotal int        `gorm:"-"`
}
