package controller

import (
	"fmt"
	"strconv"
	"wcgblog/model"
	"wcgblog/service"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//  新增评论
func CommentPost(c *gin.Context) {
	var (
		err  error
		res  = gin.H{}
		post *model.Post
	)
	fmt.Println(post)
	defer writeJSON(c, res)
	session := sessions.Default(c)
	sessionUserId := session.Get(SESSION_KEY)
	userId, _ := sessionUserId.(uint)
	verifyCode := c.PostForm("verifyCode")
	captchaId := session.Get(SESSION_CAPTCHA)
	session.Delete(SESSION_CAPTCHA)
	capId := captchaId.(string)
	if !captcha.VerifyString(capId, verifyCode) {
		res["message"] = "验证码错误"
		return
	}
	postId := c.PostForm("postId")
	content := c.PostForm("content")
	if len(content) == 0 {
		res["message"] = "content cannot be empty"
		return
	}
	post, err = service.SelectPostById(postId)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	pID, err := strconv.ParseUint(postId, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	comment := &model.Comment{
		PostID:  uint(pID),
		Content: content,
		UserID:  userId,
	}
	err = service.InsertComment(comment)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

// 删除评论
func CommentDelete(c *gin.Context) {
	var (
		err error
		res = gin.H{}
		cid uint64
	)
	defer writeJSON(c, res)
	session := sessions.Default(c)
	sessionUserId := session.Get(SESSION_KEY)
	userId := sessionUserId.(uint)
	commentId := c.Param("id")
	cid, err = strconv.ParseUint(commentId, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	comment := &model.Comment{
		UserID: uint(userId),
	}
	comment.ID = uint(cid)
	err = service.DeleteCommentId(comment)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

// 获取评论 让评论已读
func CommentRead(c *gin.Context) {
	var (
		id  uint64
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	cid := c.Param("id")
	id, err = strconv.ParseUint(cid, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	comment := &model.Comment{}
	comment.ID = uint(id)
	err = service.UpdateCommentReadState(comment)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

// 获取所有的评论
func CommentReadAll(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	err = service.SetAllCommentRead()
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}
