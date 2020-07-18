package policies

import (
	"ginweibo/app/controllers"

	"github.com/gin-gonic/gin"
)

// 无权限时
func Unauthorized(c *gin.Context) {
	controllers.RenderUnauthorized(c)
}
