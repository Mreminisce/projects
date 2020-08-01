package model

// 分页 数据表
type Page struct {
	BaseModel
	Title       string `gorm:"default:null comment '标题'"`  //标题
	Body        string `gorm:"default:null comment '主题'"`  // 主题
	View        int    `gorm:"default:0 comment '查看计数'"`   // 查看计数
	IsPublished bool   `gorm:"default:'0' comment '发表或不'"` // 发表或不
}
