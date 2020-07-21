package middleware

import (
	"ginweibo/middleware/auth"

	"github.com/gin-gonic/gin"
)

// 从 session 中获取 user model 的 middleware
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.SaveCurrentUserToContext(c)
		c.Next()
	}
}
