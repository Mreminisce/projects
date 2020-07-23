package auth

import (
	"errors"
	"ginweibo/config"
	userModel "ginweibo/models/user"
	"ginweibo/utils/rand"
	"ginweibo/utils/session"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	rememberFormKey    = "remember"
	rememberCookieName = "remember_me"
	rememberMaxAge     = 88888888 // 过期时间
)

func getRememberTokenFromCookie(c *gin.Context) string {
	if cookie, err := c.Request.Cookie(rememberCookieName); err == nil {
		if v, err := url.QueryUnescape(cookie.Value); err == nil {
			return v
		}
	}
	return ""
}

func delRememberToken(c *gin.Context) {
	c.SetCookie(rememberCookieName, "", -1, "/", "", false, true)
}

// 从 session 中获取用户
func getCurrentUserFromSession(c *gin.Context) (*userModel.User, error) {
	rememberMeToken := getRememberTokenFromCookie(c)
	if rememberMeToken != "" {
		if user, err := userModel.GetByRememberToken(rememberMeToken); err == nil {
			Login(c, user)
			return user, nil
		}
		delRememberToken(c)
	}
	idStr := session.GetSession(c, config.AppConfig.AuthSessionKey)
	if idStr == "" {
		return nil, errors.New("没有获取到 session")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	user, err := userModel.Get(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 记住我，如果登录的 PostForm 中有着 remember="on" 说明开启记住我功能
func setRememberTokenInCookie(c *gin.Context, u *userModel.User) {
	rememberMe := c.PostForm(rememberFormKey) == "on"
	if !rememberMe {
		return
	}
	// 更新用户的 RememberToken
	newToken := string(rand.RandomCreateBytes(10))
	u.RememberToken = newToken
	if err := u.Update(false); err != nil {
		return
	}
	c.SetCookie(rememberCookieName, u.RememberToken, rememberMaxAge, "/", "", false, true)
}

// Login 登录
func Login(c *gin.Context, u *userModel.User) {
	session.SetSession(c, config.AppConfig.AuthSessionKey, u.GetIDstring())
	setRememberTokenInCookie(c, u)
}

// Logout 登出
func Logout(c *gin.Context) {
	session.DeleteSession(c, config.AppConfig.AuthSessionKey)
	delRememberToken(c)
}

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
