package routes

import (
	"ginweibo/controllers/followers"
	staticpage "ginweibo/controllers/home"
	"ginweibo/controllers/password"
	"ginweibo/controllers/sessions"
	"ginweibo/controllers/status"
	"ginweibo/controllers/user"
	"ginweibo/middleware/wrapper"
	"ginweibo/routes/named"

	"github.com/gin-gonic/gin"
)

func registerWeb(g *gin.Engine) {
	// static page
	{
		g.GET("/", staticpage.Home)
		// 绑定路由 path 和路由 name，之后可通过 named.G("root") 或 named.GR("root") 获取到路由 path
		// 模板文件中可通过 {{ Route "root" }} 或 {{ RelativeRoute "root" }} 获取 path
		named.Name(g, "root", "GET", "/")
		g.GET("/help", staticpage.Help)
		named.Name(g, "help", "GET", "/help")
		g.GET("/about", staticpage.About)
		named.Name(g, "about", "GET", "/about")
	}
	// user
	{
		g.GET("/signup", wrapper.Guest(user.Create))
		named.Name(g, "signup", "GET", "/signup")
		g.GET("/signup/confirm/:token", wrapper.Guest(user.ConfirmEmail))
		// 带参路由绑定，可通过 named.G("signup.confirm", token) 或 named.GR("signup.confirm", token) 获取 path
		// 模板文件中可通过 {{ Route "signup.confirm" .token }} 或 {{ RelativeRoute "signup.confirm" .token }} 获取 path
		named.Name(g, "signup.confirm", "GET", "/signup/confirm/:token")
		userRouter := g.Group("/users")
		{
			userRouter.GET("/create", wrapper.Guest(user.Create))
			named.Name(userRouter, "users.create", "GET", "/create")
			userRouter.POST("", wrapper.Guest(user.Store))
			named.Name(userRouter, "users.store", "POST", "")
			userRouter.GET("", wrapper.Auth(user.Index)) // 用户列表
			named.Name(userRouter, "users.index", "GET", "")
			userRouter.GET("/show/:id", wrapper.Auth(user.Show)) // 展示具体用户
			named.Name(userRouter, "users.show", "GET", "/show/:id")
			userRouter.GET("/edit/:id", wrapper.Auth(user.Edit))
			named.Name(userRouter, "users.edit", "GET", "/edit/:id")
			userRouter.POST("/update/:id", wrapper.Auth(user.Update))
			named.Name(userRouter, "users.update", "POST", "/update/:id")
			userRouter.POST("/destroy/:id", wrapper.Auth(user.Destroy))
			named.Name(userRouter, "users.destroy", "POST", "/destroy/:id")
			userRouter.GET("/followings/:id", wrapper.Auth(user.Followings))
			named.Name(userRouter, "users.followings", "GET", "/followings/:id")
			userRouter.GET("/followers/:id", wrapper.Auth(user.Followers))
			named.Name(userRouter, "users.followers", "GET", "/followers/:id")
			userRouter.POST("/followers/store/:id", wrapper.Auth(followers.Store))
			named.Name(userRouter, "followers.store", "POST", "/followers/store/:id")
			userRouter.POST("/followers/destroy/:id", wrapper.Auth(followers.Destroy))
			named.Name(userRouter, "followers.destroy", "POST", "/followers/destroy/:id")
		}
	}
	// sessions
	{
		g.GET("/login", wrapper.Guest(sessions.Create))
		named.Name(g, "login.create", "GET", "/login")
		g.POST("/login", wrapper.Guest(sessions.Store))
		named.Name(g, "login.store", "POST", "/login")
		g.POST("/logout", sessions.Destroy)
		named.Name(g, "login.destroy", "POST", "/logout")
		named.Name(g, "logout", "POST", "/logout")
	}
	// password
	passwordRouter := g.Group("/password")
	{
		passwordRouter.GET("/reset", wrapper.Guest(password.ShowLinkRequestsForm))
		named.Name(passwordRouter, "password.request", "GET", "/reset")
		passwordRouter.POST("/email", wrapper.Guest(password.SendResetLinkEmail))
		named.Name(passwordRouter, "password.email", "POST", "/email")
		passwordRouter.GET("/reset/:token", wrapper.Guest(password.ShowResetForm))
		named.Name(passwordRouter, "password.reset", "GET", "/reset/:token")
		passwordRouter.POST("/reset", wrapper.Guest(password.Reset))
		named.Name(passwordRouter, "password.update", "POST", "/reset")
	}
	// statuses
	statusRouter := g.Group("/statuses")
	{
		statusRouter.POST("", wrapper.Auth(status.Store))
		named.Name(statusRouter, "statuses.store", "POST", "")
		statusRouter.POST("/destroy/:id", wrapper.Auth(status.Destroy))
		named.Name(statusRouter, "statuses.destroy", "POST", "/destroy/:id")
	}
}
