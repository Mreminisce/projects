package controller

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"wcgblog/model"
	"wcgblog/service"

	"github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

// 博客首页
func PostIndex(c *gin.Context) {
	posts, _ := service.ListAllPost("")
	user, _ := c.Get(CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "admin/post.html", gin.H{
		"posts":    posts,
		"Active":   "posts",
		"user":     user,
		"comments": service.MustListUnreadComment(),
	})
}

func PostNew(c *gin.Context) {
	c.HTML(http.StatusOK, "post/new.html", nil)
}

// 新增博客
func PostCreate(c *gin.Context) {
	tags := c.PostForm("tags")
	title := c.PostForm("title")
	body := c.PostForm("body")
	isPublished := c.PostForm("isPublished")
	published := "on" == isPublished
	post := &model.Post{
		Title:       title,
		Body:        body,
		IsPublished: published,
	}
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	err := service.PostInsert(post)
	if err != nil {
		c.HTML(http.StatusOK, "post/new.html", gin.H{
			"post":    post,
			"message": err.Error(),
		})
		return
	}
	//	 新增标签中间表
	if len(tags) > 0 {
		arr := strings.Split(tags, ",")
		for _, tag := range arr {
			tagId, err := strconv.ParseUint(tag, 10, 64)
			if err != nil {
				continue
			}
			pt := &model.PostTag{
				PostId: post.ID,
				TagId:  uint(tagId),
			}
			pt.CreatedAt = time.Now()
			pt.UpdatedAt = time.Now()
			err = service.TagPostInsert(pt)
			if err != nil {
				seelog.Debug(err.Error())
			}
		}
	}
	//	 重定向
	c.Redirect(http.StatusMovedPermanently, "/admin/post")
}

// 编辑博客
func PostEdit(c *gin.Context) {
	id := c.Param("id")
	post, err := service.SelectPostById(id)
	if err != nil {
		Handle404(c)
		return
	}
	post.Tags, _ = service.ListTagPostById(id)
	c.HTML(http.StatusOK, "post/modify.html", gin.H{
		"post": post,
	})
}

// 修改博客
func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	tags := c.PostForm("tags")
	title := c.PostForm("title")
	body := c.PostForm("body")
	isPublished := c.PostForm("isPublished")
	published := "on" == isPublished
	PID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		Handle404(c)
		return
	}
	post := &model.Post{
		Title:       title,
		Body:        body,
		IsPublished: published,
	}
	post.UpdatedAt = time.Now()
	err = service.UpdatePost(post, id)
	if err != nil {
		c.HTML(http.StatusOK, "post/modify.html", gin.H{
			"post":    post,
			"message": err.Error(),
		})
		return
	}
	//	 删除 tag
	service.DeletePostTagByPostId(PID)
	//	 添加tag
	if len(tags) > 0 {
		arr := strings.Split(tags, ",")
		for _, tag := range arr {
			tagId, err := strconv.ParseUint(tag, 10, 64)
			if err != nil {
				continue
			}
			pt := &model.PostTag{
				PostId: uint(PID),
				TagId:  uint(tagId),
			}
			pt.CreatedAt = time.Now()
			pt.UpdatedAt = time.Now()
			err = service.TagPostInsert(pt)
			if err != nil {
				seelog.Error(err.Error())
			}
		}
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/post")
}

// 设置文章是否公开
func PublishPost(c *gin.Context) {
	var (
		err  error
		res  = gin.H{}
		post *model.Post
	)
	// 最后写入json
	defer writeJSON(c, res)
	id := c.Param("id")
	post, err = service.SelectPostById(id)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	post.IsPublished = !post.IsPublished
	err = service.UpdatePost(post, id)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

// 删除博客
func DeletePost(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	id := c.Param("id")
	PID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	post := &model.Post{}
	post.ID = uint(PID)
	err = service.DeletePost(post)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	err = service.DeletePostTagByPostId(PID)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

// 查看博客详情
func GetPost(c *gin.Context) {
	var (
		res = gin.H{}
	)
	id := c.Param("id")
	post, err := service.SelectPostById(id)
	if err != nil || !post.IsPublished {
		Handle404(c)
		return
	}
	post.View++
	err = service.UpdatePostView(post)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	post.Tags, _ = service.ListTagPostById(id)
	post.Comments, err = service.ListCommentByPostID(id)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	user, _ := c.Get(CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "post/display.html", gin.H{
		"post": post,
		"user": user,
	})
}
