package service

import (
	"fmt"
	"strconv"
	"wcgblog/model"
)

func MustListTag() []*model.Tag {
	tags, _ := LisrTag()
	return tags
}

func LisrTag() ([]*model.Tag, error) {
	var tags []*model.Tag
	rows, err := db.Raw("select t.* ,count(*) total from tags t inner join post_tags pt on t.id = pt.tag_id inner join posts p on pt.post_id =p.id where p.is_published =? "+
		"group by pt.tag_id", true).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag model.Tag
		db.ScanRows(rows, &tag)
		tags = append(tags, &tag)
	}
	return tags, nil
}

func CountTag() int {
	var count int
	db.Model(&model.Tag{}).Count(&count)
	return count
}

// 新增标签博客中间表
func TagPostInsert(tag *model.PostTag) error {
	err := db.Create(tag).Error
	return err
}

// 删除关联博客的数据
func DeletePostTagByPostId(id interface{}) error {
	return db.Delete(&model.PostTag{}, "post_id =?", id).Error
}

// 新增标签
func InsertTag(tag *model.Tag) error {
	return db.FirstOrCreate(tag, "name =?", tag.Name).Error
}

// 获取数量
func CountPostByTag(tag string) (count int, err error) {
	var tagId uint64
	if len(tag) > 0 {
		tagId, err = strconv.ParseUint(tag, 10, 64)
		if err != nil {
			return
		}
		err = db.Raw("select count(*) from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ? and p.is_published = ?", tagId, true).Row().Scan(&count)
	} else {
		err = db.Raw("select count(*) from posts p where p.is_published = ?", true).Row().Scan(&count)
	}
	if err != nil {
		fmt.Println("出现查询错误:", err)
		return 0, err
	}
	return
}

func ListTagByPostId(id string) ([]*model.Tag, error) {
	var tags []*model.Tag
	pid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	rows, err := db.Raw("select t.* from tags t inner join post_tags pt on t.id = pt.tag_id where pt.post_id = ?", uint(pid)).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tag model.Tag
		db.ScanRows(rows, &tag)
		tags = append(tags, &tag)
	}
	return tags, nil
}
