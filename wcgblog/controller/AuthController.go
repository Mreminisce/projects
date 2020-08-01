package controller

import (
	"net/http"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 获取认证方式
func AuthGet(c *gin.Context) {
	authType := c.Param("authType")
	session := sessions.Default(c)
	session.Save()
	authrl := "/signin"
	switch authType {
	case "weibo":
	case "qq":
	case "wechat":
	case "oschina":
	default:
	}
	c.Redirect(http.StatusFound, authrl)
}

// 验证码
func CaptchaGet(context *gin.Context) {
	session := sessions.Default(context)
	id := captcha.NewLen(4)
	session.Delete(SESSION_CAPTCHA)
	session.Set(SESSION_CAPTCHA, id)
	session.Save()
	captcha.WriteImage(context.Writer, id, 100, 40)
}
