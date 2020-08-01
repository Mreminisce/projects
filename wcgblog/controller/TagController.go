package controller

import (
	"math"
	"net/http"
	"strconv"
	"wcgblog/model"
	"wcgblog/service"
	"wcgblog/system"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

// 新增标签
func TagCreate(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	name := c.PostForm("value")
	tag := &model.Tag{Name: name}
	err = service.InsertTag(tag)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
	res["data"] = tag
}

// 获取标签信息
func GetTag(c *gin.Context) {
	var (
		tagName   string
		page      string
		pageIndex int
		pageSize  = system.GetConfiguration().PageSize
		total     int
		err       error
		polict    *bluemonday.Policy
		posts     []*model.Post
	)
	tagName = c.Param("tag")
	page = c.Query("page")
	pageIndex, _ = strconv.Atoi(page)
	if pageIndex <= 0 {
		pageIndex = 1
	}
	posts, err = service.ListPublishedPost(tagName, pageIndex, pageSize)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	total, err = service.CountPostByTag(tagName)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	polict = bluemonday.StrictPolicy()
	for _, post := range posts {
		post.Tags, _ = service.ListTagPostById(strconv.FormatUint(uint64(post.ID), 10))
		post.Body = polict.Sanitize(string(blackfriday.Run([]byte(post.Body))))
	}
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"posts":           posts,
		"tags":            service.MustListTag(),
		"links":           service.MustListLinks(),
		"pageIndex":       pageIndex,
		"totalPage":       int(math.Ceil(float64(total) / float64(pageSize))),
		"maxReadPosts":    service.MustListMaxReadPost(),
		"maxCommentPosts": service.MustListMaxCommentPost(),
	})
}
