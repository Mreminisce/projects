package routes

import (
	"ginweibo/controllers"
	"ginweibo/middleware"

	"github.com/gin-gonic/gin"
	ginSessions "github.com/tommy351/gin-sessions"
)

var (
	sessionKeyPairs  = []byte("secret123")
	sessionStoreName = "my_session"
)

// 注册路由和中间件
func Register(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(gin.Logger())
	store := ginSessions.NewCookieStore(sessionKeyPairs)
	store.Options(ginSessions.Options{
		HttpOnly: true,
		Path:     "/",
		MaxAge:   86400 * 30,
	})
	g.Use(ginSessions.Middleware(sessionStoreName, store))
	g.Use(middleware.Csrf())     // csrf
	g.Use(middleware.OldValue()) // 记忆上次表单提交的内容，消费即消失
	g.Use(middleware.GetUser())  // 从 session 中获取用户
	g.NoRoute(func(c *gin.Context) {
		controllers.Render404(c)
	})
	registerWeb(g)
	return g
}
