package policies

import (
	"ginweibo/controllers"
	statusModel "ginweibo/models/status"
	userModel "ginweibo/models/user"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// 无权限时
func Unauthorized(c *gin.Context) {
	controllers.RenderUnauthorized(c)
}

// 是否有删除微博的权限
func StatusPolicyDestroy(c *gin.Context, currentUser *userModel.User, status *statusModel.Status) bool {
	if currentUser.ID != status.UserID {
		log.Infof("%s 没有权限删除微博 (ID: %d)", currentUser.Name, status.UserID)
		Unauthorized(c)
		return false
	}
	return true
}

// 是否有更新目标 user 的权限
func UserPolicyUpdate(c *gin.Context, currentUser *userModel.User, targetUserID int) bool {
	if currentUser.ID != uint(targetUserID) {
		Unauthorized(c)
		return false
	}
	return true
}

// 是否有删除用户的权限 (只有当前用户拥有管理员权限且删除的用户不是自己时)
func UserPolicyDestroy(c *gin.Context, currentUser *userModel.User, targetUserID int) bool {
	if currentUser.ID == uint(targetUserID) || !currentUser.IsAdminRole() {
		Unauthorized(c)
		return false
	}
	return true
}

// 是否有关注用户的权限
func UserPolicyFollow(c *gin.Context, currentUser *userModel.User, targetUserID int) bool {
	if currentUser.ID == uint(targetUserID) {
		Unauthorized(c)
		return false
	}
	return true
}
