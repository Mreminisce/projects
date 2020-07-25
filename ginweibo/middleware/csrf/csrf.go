package csrf

import (
	"ginweibo/config"
	"ginweibo/controllers"
	"ginweibo/utils/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

// cookie 中获取 csrf token，没有则设置
func getCsrfTokenFromCookie(c *gin.Context) (token string) {
	keyName := config.AppConfig.CsrfParamName
	if s, err := c.Request.Cookie(keyName); err == nil {
		token = s.Value
	}
	if token == "" {
		token = string(rand.RandomCreateBytes(32))
		c.SetCookie(keyName, token, 0, "/", "", false, false)
	}
	c.Keys[keyName] = token
	return token
}

// 从 params 或 headers 中获取 csrf token
func getCsrfTokenFromParamsOrHeader(c *gin.Context) (token string) {
	req := c.Request
	if req.Form == nil {
		req.ParseForm()
	}
	token = req.FormValue(config.AppConfig.CsrfParamName)
	if token == "" {
		token = req.Header.Get(config.AppConfig.CsrfHeaderName)
	}
	return token
}

func Csrf() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.AppConfig.EnableCsrf {
			csrfToken := getCsrfTokenFromCookie(c)
			if c.Request.Method == http.MethodPost {
				paramCsrfToken := getCsrfTokenFromParamsOrHeader(c)
				if paramCsrfToken == "" || paramCsrfToken != csrfToken {
					controllers.Render403(c, "您的 Session 已过期，刷新后再试一次。")
					c.Abort()
					return
				}
			}
		}
		c.Next()
	}
}
