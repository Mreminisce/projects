package controller

import (
	"net/http"
	"strconv"
	"wcgblog/model"
	"wcgblog/service"

	"github.com/gin-gonic/gin"
)

// 链接首页
func LinkIndex(c *gin.Context) {
	links, _ := service.ListLinks()
	user, _ := c.Get(CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "admin/link.html", gin.H{
		"links":    links,
		"user":     user,
		"comments": service.MustListUnreadComment(),
	})
}

// 获取链接
func LinkGet(c *gin.Context) {
	id := c.Param("id")
	link, err := service.GetLinkById(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	link.View++
	err = service.UpdateLink(link)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusFound, link.Url)
}

// 修改链接
func LinkCreate(c *gin.Context) {
	var (
		err   error
		res   = gin.H{}
		_sort int64
	)
	defer writeJSON(c, res)
	name := c.PostForm("name")
	url := c.PostForm("url")
	sort := c.PostForm("sort")
	if len(name) == 0 || len(url) == 0 {
		res["message"] = "错误参数"
		return
	}
	_sort, err = strconv.ParseInt(sort, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	link := &model.Link{
		Name: name,
		Url:  url,
		Sort: int(_sort),
	}
	err = service.InsertLink(link)
	if err != nil {
		res["message"] = "新增失败"
		return
	}
	res["succeed"] = true
}

// 修改连接
func LinkUpdate(c *gin.Context) {
	var (
		_id   uint64
		_sort int64
		err   error
		res   = gin.H{}
	)
	defer writeJSON(c, res)
	id := c.Param("id")
	name := c.PostForm("name")
	url := c.PostForm("url")
	sort := c.PostForm("sort")
	if len(name) == 0 || len(url) == 0 || len(sort) == 0 {
		res["message"] = "参数不完整"
		return
	}
	_id, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	_sort, err = strconv.ParseInt(sort, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	link := &model.Link{
		Name: name,
		Url:  url,
		Sort: int(_sort),
	}
	link.ID = uint(_id)
	err = service.UpdateLink(link)
	if err != nil {
		res["message"] = "修改失败"
		return
	}
	res["succeed"] = true
}

// 删除链接
func LinkDelete(c *gin.Context) {
	var (
		err error
		_id uint64
		res = gin.H{}
	)
	defer writeJSON(c, res)
	id := c.Param("id")
	_id, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	link := &model.Link{}
	link.ID = uint(_id)
	err = service.DeleteLink(link)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}
