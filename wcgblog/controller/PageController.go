package controller

import (
	"net/http"
	"strconv"
	"wcgblog/model"
	"wcgblog/service"

	"github.com/gin-gonic/gin"
)

// 首页
func PageIndex(c *gin.Context) {
	pages, err := service.ListAllPage()
	if err != nil {
		return
	}
	user, _ := c.Get(CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "admin/page.html", gin.H{
		"pages":    pages,
		"user":     user,
		"comments": service.MustListUnreadComment(),
	})
}

// 新建页面
func PageNew(c *gin.Context) {
	c.HTML(http.StatusOK, "page/new.html", nil)
}

// 创建页面
func PageCreate(c *gin.Context) {
	var (
		err error
	)
	title := c.PostForm("title")
	body := c.PostForm("body")
	isPublished := c.PostForm("isPublished")
	published := "on" == isPublished
	page := &model.Page{
		Title:       title,
		Body:        body,
		IsPublished: published,
	}
	err = service.PageInsert(page)
	if err != nil {
		c.HTML(http.StatusOK, "page/new.html", gin.H{
			"message": err.Error(),
			"page":    page,
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/page")
}

// 编辑页面
func PageEdit(c *gin.Context) {
	id := c.Param("id")
	page, err := service.SelectPageById(id)
	if err != nil {
		Handle404(c)
	}
	c.HTML(http.StatusOK, "page/modify.html", gin.H{"page": page})
}

// 修改页面
func UpdatePage(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	body := c.PostForm("body")
	isPublished := c.PostForm("isPublished")
	publish := "on" == isPublished
	PID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	page := &model.Page{
		Title:       title,
		Body:        body,
		IsPublished: publish,
	}
	page.ID = uint(PID)
	err = service.UpdatePage(page)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/admin/page")
}

// 修改状态
func PagePublish(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	id := c.Param("id")
	page, err := service.SelectPageById(id)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	page.IsPublished = !page.IsPublished
	err = service.UpdatePage(page)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

// 删除页面
func DeletePage(c *gin.Context) {
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
	page := &model.Page{}
	page.ID = uint(PID)
	err = service.DeletePage(page)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

func PageGet(c *gin.Context) {
	id := c.Param("id")
	page, err := service.SelectPageById(id)
	if err != nil || !page.IsPublished {
		Handle404(c)
		return
	}
	page.View++
	service.UpdatePageViem(page)
	c.HTML(http.StatusOK, "page/display.html", gin.H{
		"page": page,
	})
}
