package service

import (
	"fmt"
	"strconv"
	"wcgblog/model"
)

func CountComment() int {
	var count int
	db.Model(&model.Comment{}).Count(&count)
	return count
}

func ListUnreadComment() ([]*model.Comment, error) {
	var comments []*model.Comment
	err := db.Where("read_state = ?", false).Order("created_at desc").Find(&comments).Error
	if err != nil {
		fmt.Println("出现查询错误:", err)
		return nil, err
	}
	return comments, err
}

func MustListUnreadComment() []*model.Comment {
	comments, _ := ListUnreadComment()
	return comments
}

// 通过博客id  获取博客的评论
func ListCommentByPostID(postId string) ([]*model.Comment, error) {
	pid, err := strconv.ParseUint(postId, 10, 64)
	if err != nil {
		return nil, err
	}
	var comments []*model.Comment
	rows, err := db.Raw("select c.*,u.avatar_url  from comments c "+
		"inner join users u on c.user_id = u.id where c.post_id = ? order by created_at desc ", uint(pid)).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment model.Comment
		// 地址赋值过去
		db.ScanRows(rows, &comment)
		comments = append(comments, &comment)
	}
	return comments, err
}

// 新增评论
func InsertComment(comment *model.Comment) error {
	return db.Create(&comment).Error
}

// 删除一条评论
func DeleteCommentId(comment *model.Comment) error {
	return db.Delete(comment, "user_id = ? and id =?", comment.UserID, comment.ID).Error
}

// 修改阅读状态
func UpdateCommentReadState(comment *model.Comment) error {
	return db.Model(&comment).Updates(map[string]interface{}{
		"read_state": true,
	}).Error
}

// 获取所有的已经阅读的信息
func SetAllCommentRead() error {
	return db.Model(&model.Comment{}).Where("read_state =?", false).Updates(map[string]interface{}{
		"read_state": true,
	}).Error
}
