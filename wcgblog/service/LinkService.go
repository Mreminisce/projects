package service

import (
	"fmt"
	"wcgblog/model"
)

func MustListLinks() []*model.Link {
	links, _ := ListLinks()
	return links
}
func ListLinks() ([]*model.Link, error) {
	var links []*model.Link
	err := db.Order("sort desc").Find(&links).Error
	if err != nil {
		fmt.Println("出现查询错误:", err)
		return nil, err
	}
	return links, err
}

//  通过id获取链接
func GetLinkById(id interface{}) (*model.Link, error) {
	var link model.Link
	err := db.FirstOrCreate(&link, "id =?", id).Error
	if err != nil {
		fmt.Println("出现查询错误:", err)
		return nil, err
	}
	return &link, err
}

// 修改链接
func UpdateLink(link *model.Link) error {
	return db.Save(&link).Error
}

// 新增链接
func InsertLink(link *model.Link) error {
	return db.FirstOrCreate(link, "url = ?", link.Url).Error
}

// 删除链接
func DeleteLink(link *model.Link) error {
	return db.Delete(link).Where("id = ?", link.ID).Error
}
