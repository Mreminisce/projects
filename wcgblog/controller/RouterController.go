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
	g := gin.Default()
	// 配置文件解析
	setTemplate(g)
	// session 初始化
	setSessions(g)
	// 填充常用信息
	g.Use(SharedData())
	// 定时任务，每天获取一次网页数据地图
	gocron.Every(1).Day().Do(CreateXMLSitemap)
	gocron.Start()
	g.Static("/static", filepath.Join(getCurrentDirectory(), "./static"))
	g.NoRoute(Handle404)
	// 首页
	g.GET("/", IndexGet)
	g.GET("/index", IndexGet)
	// 用户相关接口
	if system.GetConfiguration().SignupEnabled {
		g.GET("/signup", SignupGet)
		g.POST("/signup", SignupPost)
	}
	g.GET("/signin", SigninGet)
	g.POST("/signin", SigninPost)
	g.GET("/logout", LogoutGet)
	// 认证
	g.GET("/auth/:authType", AuthGet)
	// 验证码
	g.GET("/captcha", CaptchaGet)
	// 获取博客详情
	g.GET("/post/:id", GetPost)
	// 获取页面详情
	g.GET("/page/:id", PageGet)
	// 获取标签
	g.GET("/tag/:tag", GetTag)
	// 获取链接
	g.GET("/link/:id", LinkGet)
	visitor := g.Group("/visitor")
	visitor.Use(AuthRequired())
	{
		// 新增评论
		visitor.POST("/new_comment", CommentPost)
		//删除评论
		visitor.POST("/comment/:id/delete", CommentDelete)
	}
	admin := g.Group("/admin")
	admin.Use(AdminScopeRequired())
	{
		// 首页
		admin.GET("/index", AdminIndex)
		//	 用户首页
		admin.GET("/user", UserIndex)
		admin.GET("/user/:id/lock", UserLock)
		//	 编写博客与修改删除博客
		admin.GET("/post", PostIndex)
		admin.GET("/new_post", PostNew)
		admin.POST("/new_post", PostCreate)
		admin.GET("/post/:id/edit", PostEdit)
		admin.POST("/post/:id/edit", UpdatePost)
		admin.POST("/post/:id/publish", PublishPost)
		admin.POST("/post/:id/delete", DeletePost)
		//	 页面管理
		admin.GET("/page", PageIndex)
		admin.GET("/new_page", PageNew)
		admin.POST("/new_page", PageCreate)
		admin.GET("/page/:id/edit", PageEdit)
		admin.POST("/page/:id/edit", UpdatePage)
		admin.POST("/page/:id/publish", PagePublish)
		admin.POST("/page/:id/delete", DeletePage)
		//	读取评论
		admin.POST("/comment/:id", CommentRead)
		admin.POST("/read_all", CommentReadAll)
		// 标签
		admin.POST("/new_tag", TagCreate)
		//	 链接
		admin.GET("/link", LinkIndex)
		admin.POST("/new_link", LinkCreate)
		admin.POST("/link/:id/edit", LinkUpdate)
		admin.POST("/link/:id/delete", LinkDelete)
		//	 用户信息
		admin.GET("/profile", ProfileGet)
		admin.POST("/profile", ProfileUpdate)
		admin.POST("/profile/email/bind", BindEmail)
		admin.POST("/profile/email/unbind", UnbinEmail)
	}
	return g
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
func setSessions(g *gin.Engine) {
	config := system.GetConfiguration()
	store := cookie.NewStore([]byte(config.SessionSecret))
	//	 MaxAge 时间
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	g.Use(sessions.Sessions("gin-session", store))
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
		seelog.Warnf("User not admin to visit %s", c.Request.RequestURI)
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
		seelog.Warnf("User not admin to visit %s", c.Request.RequestURI)
		c.HTML(http.StatusForbidden, "errors/error.html", gin.H{
			"message": "Forbidden !",
		})
		c.Abort()
	}
}
