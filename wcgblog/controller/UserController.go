package controller

import (
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"net/http"
	"wcgblog/model"
	"wcgblog/service"
	"wcgblog/utils"
)

// 获取登录
func SigninGet(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/signin.html", nil)
}

// 获取注册
func SignupGet(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/signup.html", nil)
}

// 退出
func LogoutGet(c *gin.Context) {
	// 获取session会话
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	// 重定向
	c.Redirect(http.StatusSeeOther, "/signin")
}

// 注册
func SignupPost(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	// 获取表单数据
	defer writeJSON(c, res)
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := &model.User{
		Email:    email,
		Password: password,
	}
	if len(user.Email) == 0 || len(user.Password) == 0 {
		// 另一种json格式的返回
		res["message"] = "email or password cannot be null"
		return
	}
	user.Password = utils.Md5(user.Password + user.Email)
	err = service.UserInsert(user)
	if err != nil {
		res["message"] = "email already exists"
		return
	}
	res["succeed"] = true
	c.Redirect(http.StatusSeeOther, "/signin")
}

// 登录
func SigninPost(c *gin.Context) {
	var (
		err  error
		user *model.User
	)
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.HTML(http.StatusOK, "auth/signin.html", gin.H{
			"message": "username or password cannot be null",
		})
		return
	}
	user, err = service.GetUserByUsername(username)
	if err != nil || user.Password != utils.Md5(password+username) {
		c.HTML(http.StatusOK, "auth/signin.html", gin.H{
			"message": "密码或者用户名错误",
		})
		return
	}
	if user.LockState {
		c.HTML(http.StatusOK, "auth/signin.html", gin.H{
			"message": "你的账号已经被锁定",
		})
		return
	}
	// 获取session
	session := sessions.Default(c)
	session.Clear()
	session.Set(SESSION_KEY, user.ID)
	session.Save()
	// 判断是否是管理员
	// mysql 中0表示false 1表示true
	if user.IsAdmin {
		c.Redirect(http.StatusMovedPermanently, "/admin/index")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func writeJSON(ctx *gin.Context, h gin.H) {
	if _, ok := h["succeed"]; !ok {
		h["succeed"] = false
	}
	ctx.JSON(http.StatusOK, h)
}

// 用户首页
func UserIndex(c *gin.Context) {
	users, _ := service.ListUsers()
	user, _ := c.Get(CONTEXT_USER_KEY)
	c.HTML(http.StatusOK, "admin/user.html", gin.H{
		"users":    users,
		"user":     user,
		"comments": service.MustListUnreadComment(),
	})
}

// 用户锁
func UserLock(c *gin.Context) {
	var (
		err  error
		_id  uint64
		res  = gin.H{}
		user *model.User
	)
	defer writeJSON(c, res)
	id := c.Param("id")
	// 将字符串解析为整数
	_id, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	user, err = service.GetUser(uint(_id))
	if err != nil {
		res["message"] = err.Error()
		return
	}
	user.LockState = !user.LockState
	err = service.Lock(user)
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

// 获取用户信息以及评论
func ProfileGet(c *gin.Context) {
	sessionUser, exists := c.Get(CONTEXT_USER_KEY)
	if exists {
		c.HTML(http.StatusOK, "admin/profile.html", gin.H{
			"user":     sessionUser,
			"comments": service.MustListUnreadComment(),
		})
	}
}

// 修改信息
func ProfileUpdate(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	avatarUrl := c.PostForm("avatarUrl")
	users, exists := c.Get(CONTEXT_USER_KEY)
	if !exists {
		res["message"] = "没有获取到用户信息"
		return
	}
	user := new(model.User)
	// 接口类型转换为实体类型
	user, ok := users.(*model.User)
	if !ok {
		res["message"] = "用户信息转换失败"
		return
	}
	err = service.UpdateProfile(user, avatarUrl)
	if err != nil {
		res["message"] = "修改信息失败:" + err.Error()
		return
	}
	res["succeed"] = true
	res["user"] = model.User{AvatarUrl: avatarUrl}
}

// 绑定邮箱
func BindEmail(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	email := c.PostForm("email")
	sessionUser, exists := c.Get(CONTEXT_USER_KEY)
	if !exists {
		res["message"] = "没有获取到用户信息"
		return
	}
	user, ok := sessionUser.(*model.User)
	if !ok {
		res["message"] = "用户信息转换失败"
		return
	}
	if len(user.Email) > 0 {
		res["message"] = "请传入邮箱"
		return
	}
	_, err = service.GetUserByUsername(email)
	if err != nil {
		res["message"] = "用户没有邮箱信息"
		return
	}
	err = service.UpdateUserEmail(user, email)

	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}

// 解除邮箱绑定
func UnbinEmail(c *gin.Context) {
	var (
		err error
		res = gin.H{}
	)
	defer writeJSON(c, res)
	sessionUser, exists := c.Get(CONTEXT_USER_KEY)
	if !exists {
		res["message"] = "没有获取到用户信息"
		return
	}
	user, ok := sessionUser.(*model.User)
	if !ok {
		res["message"] = "用户转换失败"
		return
	}
	if user.Email == "" {
		res["message"] = "没有邮箱信息"
		return
	}
	err = service.UpdateUserEmail(user, "")
	if err != nil {
		res["message"] = err.Error()
		return
	}
	res["succeed"] = true
}
