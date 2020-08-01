package service

import (
	"fmt"
	"strconv"
	"wcgblog/model"
)

func CountPage() int {
	var count int
	db.Model(&model.Page{}).Count(&count)
	return count
}

func ListPubilshePage() ([]*model.Page, error) {
	return _listPage(true)
}

func ListAllPage() ([]*model.Page, error) {
	return _listPage(false)
}

func _listPage(pubilished bool) ([]*model.Page, error) {
	var pages []*model.Page
	var err error
	if pubilished {
		err = db.Where("is_published = ?", true).Find(&pages).Error
	} else {
		err = db.Find(&pages).Error
	}
	if err != nil {
		fmt.Println("出现查询错误:", err)
		return nil, err
	}
	return pages, err
}

func PageInsert(page *model.Page) error {
	return db.Create(&page).Error
}

// 修改页面
func UpdatePage(page *model.Page) error {
	return db.Model(page).Updates(map[string]interface{}{
		"title":        page.Title,
		"body":         page.Body,
		"is_published": page.IsPublished,
	}).Error
}

// 通过id获取信息
func SelectPageById(id string) (*model.Page, error) {
	Pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	var page model.Page
	err = db.First(&page, "id = ?", Pid).Error
	if err != nil {
		fmt.Println("出现查询错误:", err)
		return nil, err
	}
	return &page, err
}

// 删除
func DeletePage(page *model.Page) error {
	return db.Delete(page).Error
}

// 更新页面读取数量
func UpdatePageViem(page *model.Page) error {
	return db.Model(page).Updates(map[string]interface{}{
		"view": page.View,
	}).Error
}
