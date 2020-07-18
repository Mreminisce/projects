package wrapper

import (
	"ginweibo/app/auth"
	"ginweibo/app/controllers"
	"ginweibo/pkg/flash"

	"github.com/gin-gonic/gin"
)

// 非登录用户访问
func Guest(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 用户已经登录了则跳转到 root page
		currentUser, err := auth.GetCurrentUserFromContext(c)
		if currentUser != nil || err == nil {
			flash.NewInfoFlash(c, "您已登录，无需再次操作。")
			controllers.RedirectRouter(c, "root")
			return
		}
		handler(c)
	}
}
