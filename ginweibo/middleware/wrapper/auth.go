package wrapper

import (
	"ginweibo/controllers"
	"ginweibo/middleware/auth"
	"ginweibo/middleware/flash"
	userModel "ginweibo/models/user"

	"github.com/gin-gonic/gin"
)

type (
	AuthHandlerFunc = func(*gin.Context, *userModel.User)
)

func Auth(handler AuthHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 用户未登录则跳转到登录页
		currentUser, err := auth.GetCurrentUserFromContext(c)
		if currentUser == nil || err != nil {
			controllers.RedirectToLoginPage(c)
			return
		}
		handler(c, currentUser)
	}
}

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
