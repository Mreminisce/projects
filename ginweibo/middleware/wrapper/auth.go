package wrapper

import (
	"ginweibo/middleware/auth"
	"ginweibo/controllers"
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
