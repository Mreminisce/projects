package service

import (
	"database/sql"
	"fmt"
	"html/template"
	"strconv"
	"wcgblog/model"

	"github.com/cihub/seelog"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func ListPublishedPost(tag string, pageIndex, pageSize int) ([]*model.Post, error) {
	return _listPost(tag, true, pageIndex, pageSize)
}

func ListAllPost(tag string) ([]*model.Post, error) {
	return _listPost(tag, false, 0, 0)
}

// 获取符合要求的博客 ，包括 发刊的或者未发刊的
func _listPost(tag string, published bool, pageIndex, pageSize int) ([]*model.Post, error) {
	var posts []*model.Post
	var err error
	if len(tag) > 0 {
		//	将字符串解析为整数 标签id
		tagId, err := strconv.ParseUint(tag, 10, 64)
		if err != nil {
			seelog.Critical("解析字符串错误：", err)
		}
		// 返回查询的所有结果
		var rows *sql.Rows
		// 是否发刊
		if published {
			// db.Raw 表示执行原生sql
			if pageIndex > 0 {
				rows, err = db.Raw("select p.* from posts p left join post_tags pt on p.id = pt.post_id where pt.tag_id =? and p.is_published = ? "+
					"order by created_at desc limit ? , ? ", tagId, true, (pageIndex-1)*pageSize, pageSize).Rows()
			} else {
				rows, err = db.Raw("select p.* from posts p left join post_tag pt on p.id = pt.post_id where pt.tag_id = ? and p.is_published =? order by "+
					"created_at desc", tagId, true).Rows()
			}
		} else {
			rows, err = db.Raw("select p.* from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ? order by created_at desc", tagId).Rows()
		}
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var post model.Post
			// 迭代中使用sql.Rows的scan
			// 作用是扫描一行数据到 post实体里面
			// sacnrows 输入
			db.ScanRows(rows, &post)
			// 存放到数组
			posts = append(posts, &post)
		}
	} else {
		// 发刊
		if published {
			if pageIndex > 0 {
				// find查询符合条件的所有数据
				err = db.Where("is_published = ?", true).Order("created_at desc").Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&posts).Error
			} else {
				err = db.Where("is_published = ?", true).Order("created_at desc").Find(&posts).Error
			}
		} else {
			err = db.Order("created_at desc").Find(&posts).Error
		}
	}
	if err != nil {
		fmt.Println("出现查询错误:", err)
		return nil, err
	}
	return posts, err
}

func SelectPostByTagId(tag string) (count int, err error) {
	var (
		tagId uint64
	)
	if len(tag) > 0 {
		//	将字符串转换为整数
		tagId, err = strconv.ParseUint(tag, 10, 64)
		if err != nil {
			return
		}
		// scan 输入
		err = db.Raw("select count(*) from posts p inner join post_tags pt on p.id = pt.post_id where pt.tag_id = ? and p.is_published= ?", tagId, true).Row().Scan(&count)
	} else {
		err = db.Raw("select count(*) from posts p where  p.is_published = ?", true).Row().Scan(&count)
	}
	if err != nil {
		fmt.Println("出现查询错误:", err)
		return 0, err
	}
	return
}

func MustListMaxReadPost() (posts []*model.Post) {
	posts, _ = ListMaxReadPost()
	return
}

// 获取阅读数量最多的五条博客
func ListMaxReadPost() (posts []*model.Post, err error) {
	err = db.Where("is_published = ? ", true).Order("view desc").Limit(5).Find(&posts).Error
	if err != nil {
		fmt.Println("出现查询错误:", err)
		return nil, err
	}
	return posts, err
}

func MustListMaxCommentPost() (posts []*model.Post) {
	posts, _ = ListMaxCommentPost()
	return
}

//  获取评论最多的五条博客
func ListMaxCommentPost() (posts []*model.Post, err error) {
	var (
		rows *sql.Rows
	)
	rows, err = db.Raw("SELECT	p.*,	c.total comment_total FROM	posts p " +
		" INNER JOIN ( SELECT post_id, count( 1 ) total FROM comments GROUP BY post_id ) c ON p.id = c.post_id ORDER BY	c.total DESC 	LIMIT 5 ").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post model.Post
		db.ScanRows(rows, &post)
		posts = append(posts, &post)
	}
	return
}

func CountPost() int {
	var count int
	db.Model(&model.Post{}).Count(&count)
	return count
}

// 摘要
func Excerpt(post *model.Post) template.HTML {
	// 可以进行添加图片等
	// 删除所有的html
	policy := bluemonday.StrictPolicy()
	//	 清除模板样式
	sanitize := policy.Sanitize(string(blackfriday.Run([]byte(post.Body))))
	runes := []rune(sanitize)
	if len(runes) > 300 {
		sanitize = string(runes[:300])
	}
	excerpt := template.HTML(sanitize + "...")
	return excerpt
}

// 新增博客
func PostInsert(post *model.Post) error {
	err := db.Create(post).Error
	return err
}

// 通过id获取博客
func SelectPostById(id interface{}) (*model.Post, error) {
	var post model.Post
	err := db.First(&post, "id =?", id).Error
	if err != nil {
		fmt.Println("出现查询错误:", err)
		return nil, err
	}
	return &post, err
}

// 修改博客
func UpdatePost(post *model.Post, id string) error {
	return db.Model(post).Where("id = ?", id).Updates(map[string]interface{}{
		"title":        post.Title,
		"body":         post.Body,
		"is_published": post.IsPublished,
	}).Error
}

// 删除博客
func DeletePost(post *model.Post) error {
	return db.Delete(post).Error
}

// 修改文章阅读量
func UpdatePostView(post *model.Post) error {
	return db.Model(&post).Updates(map[string]interface{}{
		"view": post.View,
	}).Error
}
