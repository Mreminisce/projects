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

// 首页
func IndexGet(c *gin.Context) {
	var (
		pageIndex int
		pageSize  = system.GetConfiguration().PageSize
		total     int
		page      string
		err       error
		posts     []*model.Post
		policy    *bluemonday.Policy
	)
	// 获取参数
	page = c.Query("page")
	// 将string类型转换为int
	pageIndex, _ = strconv.Atoi(page)
	if pageIndex <= 0 {
		pageIndex = 1
	}
	posts, err = service.ListPublishedPost("", pageIndex, pageSize)
	if err != nil {
		//	 终止请求
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	//	 获取博客数量
	total, _ = service.SelectPostByTagId("")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	policy = bluemonday.StrictPolicy()
	for _, post := range posts {
		post.Tags, _ = service.ListTagByPostId(strconv.FormatUint(uint64(post.ID), 10))
		post.Body = policy.Sanitize(string(blackfriday.Run([]byte(post.Body))))
	}
	user, _ := c.Get(CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"posts":           posts,
		"tags":            service.MustListTag(),
		"links":           service.MustListLinks(),
		"user":            user,
		"pageIndex":       pageIndex,
		"totalPage":       int(math.Ceil(float64(total) / float64(pageSize))), //取整数
		"path":            c.Request.URL.Path,
		"maxReadPosts":    service.MustListMaxReadPost(),    // 获取五条浏览最多的数据
		"maxCommentPosts": service.MustListMaxCommentPost(), // 获取五条评论最多的数据
	})
}

func AdminIndex(c *gin.Context) {
	user, _ := c.Get(CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "admin/index.html", gin.H{
		"pageCount":    service.CountPage(),
		"postCount":    service.CountPost(),
		"tagCount":     service.CountTag(),
		"commentCount": service.CountComment(),
		"user":         user,
		"comments":     service.MustListUnreadComment(),
	})
}
