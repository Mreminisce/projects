package policies

import (
	statusModel "ginweibo/app/models/status"
	userModel "ginweibo/app/models/user"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// 是否有删除微博的权限
func StatusPolicyDestroy(c *gin.Context, currentUser *userModel.User, status *statusModel.Status) bool {
	if currentUser.ID != status.UserID {
		log.Infof("%s 没有权限删除微博 (ID: %d)", currentUser.Name, status.UserID)
		Unauthorized(c)
		return false
	}
	return true
}
