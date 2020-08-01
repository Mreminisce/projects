package controller

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"wcgblog/model"
	"wcgblog/service"
	"wcgblog/system"
	"wcgblog/utils"

	"github.com/cihub/seelog"
	"github.com/claudiu/gocron"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// 绑定所有的url
func MapRoutes() *gin.Engine {
	router := gin.Default()
	// 配置文件解析
	setTemplate(router)
	// session 初始化
	setSessions(router)
	// 填充常用信息
	router.Use(SharedData())
	// 定时任务，每天获取一次网页数据地图
	gocron.Every(1).Day().Do(CreateXMLSitemap)
	gocron.Start()
	router.Static("/static", filepath.Join(getCurrentDirectory(), "./static"))
	router.NoRoute(Handle404)
	// 首页
	router.GET("/", IndexGet)
	router.GET("/index", IndexGet)
	// 用户相关接口
	if system.GetConfiguration().SignupEnabled {
		router.GET("/signup", SignupGet)
		router.POST("/signup", SignupPost)
	}
	router.GET("/signin", SigninGet)
	router.POST("/signin", SigninPost)
	router.GET("/logout", LogoutGet)
	// 认证
	router.GET("/auth/:authType", AuthGet)
	// 验证码
	router.GET("/captcha", CaptchaGet)
	// 获取博客详情
	router.GET("/post/:id", GetPost)
	// 获取页面详情
	router.GET("/page/:id", PageGet)
	// 获取标签
	router.GET("/tag/:tag", GetTag)
	// 获取链接
	router.GET("/link/:id", LinkGet)
	visitor := router.Group("/visitor")
	visitor.Use(AuthRequired())
	{
		// 新增评论
		visitor.POST("/new_comment", CommentPost)
		//删除评论
		visitor.POST("/comment/:id/delete", CommentDelete)
	}
	authorized := router.Group("/admin")
	authorized.Use(AdminScopeRequired())
	{
		// 首页
		authorized.GET("/index", AdminIndex)
		//	 用户首页
		authorized.GET("/user", UserIndex)
		authorized.GET("/user/:id/lock", UserLock)
		//	 编写博客与修改删除博客
		authorized.GET("/post", PostIndex)
		authorized.GET("/new_post", PostNew)
		authorized.POST("/new_post", PostCreate)
		authorized.GET("/post/:id/edit", PostEdit)
		authorized.POST("/post/:id/edit", UpdatePost)
		authorized.POST("/post/:id/publish", PublishPost)
		authorized.POST("/post/:id/delete", DeletePost)
		//	 页面管理
		authorized.GET("/page", PageIndex)
		authorized.GET("/new_page", PageNew)
		authorized.POST("/new_page", PageCreate)
		authorized.GET("/page/:id/edit", PageEdit)
		authorized.POST("/page/:id/edit", UpdatePage)
		authorized.POST("/page/:id/publish", PagePublish)
		authorized.POST("/page/:id/delete", DeletePage)
		//	读取评论
		authorized.POST("/comment/:id", CommentRead)
		authorized.POST("/read_all", CommentReadAll)
		// 标签
		authorized.POST("/new_tag", TagCreate)
		//	 链接
		authorized.GET("/link", LinkIndex)
		authorized.POST("/new_link", LinkCreate)
		authorized.POST("/link/:id/edit", LinkUpdate)
		authorized.POST("/link/:id/delete", LinkDelete)
		//	 用户信息
		authorized.GET("/profile", ProfileGet)
		authorized.POST("/profile", ProfileUpdate)
		authorized.POST("/profile/email/bind", BindEmail)
		authorized.POST("/profile/email/unbind", UnbinEmail)
	}
	return router
}

// 获取当前目录
func getCurrentDirectory() string {
	// 获取文件路径 在服务器上面放开上面的注释 注释下一句代码
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return ""
	file, err := os.Getwd()
	if err != nil {
		seelog.Critical(err)
	}
	return strings.Replace(file, "\\", "/", -1)
}

//engine是Gin框架最重要的数据结构，它是框架的入口，我们通过Engine对象来定义服务路由信息，组装插件，运行服务，
func setTemplate(engine *gin.Engine) {
	funcMap := template.FuncMap{
		"dateFormat": utils.DateFormat,
		"substring":  utils.Substring,
		"isOdd":      utils.IsOdd,
		"isEven":     utils.IsEven,
		"truncate":   utils.Truncate,
		"add":        utils.Add,
		"minus":      utils.Minus,
		"listtag":    utils.ListTag,
	}
	engine.SetFuncMap(funcMap)
	// filepath.join 表示将多个元素合并成一个路径，清理多余的字符
	// 注册一个路径，gin 加载模板的时候会从该目录查找
	// 参数是一个匹配字符，如 views/*/* 的意思是 模板目录有两层
	// gin 在启动时会自动把该目录的文件编译一次缓存，不用担心效率问题
	engine.LoadHTMLGlob(filepath.Join(getCurrentDirectory(), "./views/**/*"))
}

// session初始化
func setSessions(router *gin.Engine) {
	config := system.GetConfiguration()
	store := cookie.NewStore([]byte(config.SessionSecret))
	//	 MaxAge 时间
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	router.Use(sessions.Sessions("gin-session", store))
}

// 填充常用信息 如用户信息
func SharedData() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		if ID := session.Get(SESSION_KEY); ID != nil {
			user, err := service.GetUser(ID)
			if err == nil {
				context.Set(CONTEXT_USER_KEY, user)
			}
		}
		if system.GetConfiguration().SignupEnabled {
			context.Set("SignupEnabled", true)
		}
		context.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get(CONTEXT_USER_KEY); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		seelog.Warnf("User not authorized to visit %s", c.Request.RequestURI)
		c.HTML(http.StatusForbidden, "errors/error.html", gin.H{
			"message": "Forbidden !",
		})
		c.Abort()
	}
}

// AuthRequired授予经过身份验证的用户访问权限，需要SharedData中间件
func AdminScopeRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get(CONTEXT_USER_KEY); user != nil {
			if u, ok := user.(*model.User); ok && u.IsAdmin {
				c.Next()
				return
			}
		}
		seelog.Warnf("User not authorized to visit %s", c.Request.RequestURI)
		c.HTML(http.StatusForbidden, "errors/error.html", gin.H{
			"message": "Forbidden !",
		})
		c.Abort()
	}
}
