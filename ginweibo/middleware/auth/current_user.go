package auth

import (
	"errors"
	"ginweibo/config"
	userModel "ginweibo/models/user"

	"github.com/gin-gonic/gin"
)

// 保存用户数据到 context 中
func SaveCurrentUserToContext(c *gin.Context) {
	user, err := getCurrentUserFromSession(c)
	if err != nil {
		return
	}
	c.Keys[config.AppConfig.ContextCurrentUserDataKey] = user
}

// 从 session 中获取 user model 的 middleware
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		SaveCurrentUserToContext(c)
		c.Next()
	}
}

// 从 context 中获取用户模型
func GetCurrentUserFromContext(c *gin.Context) (*userModel.User, error) {
	err := errors.New("没有获取到用户数据")
	userDataFromContext := c.Keys[config.AppConfig.ContextCurrentUserDataKey]
	if userDataFromContext == nil {
		return nil, err
	}
	user, ok := userDataFromContext.(*userModel.User)
	if !ok {
		return nil, err
	}
	return user, nil
}

// 从 context 或者从数据库中获取用户模型
func GetUserFromContextOrDataBase(c *gin.Context, id int) (*userModel.User, error) {
	// 当前用户存在并且就是想要获取的那个用户
	currentUser, err := GetCurrentUserFromContext(c)
	if currentUser != nil && err == nil {
		if int(currentUser.ID) == id {
			return currentUser, nil
		}
	}
	// 获取的是其他指定 id 的用户
	otherUser, err := userModel.Get(id)
	if err != nil {
		return nil, err
	}
	return otherUser, nil
}
