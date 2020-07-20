package main

import (
	"fmt"
	"ginweibo/app/helpers"
	"ginweibo/config"
	"ginweibo/database"
	followerModel "ginweibo/models/follower"
	passwordResetModel "ginweibo/models/password_reset"
	statusModel "ginweibo/models/status"
	userModel "ginweibo/models/user"
	"ginweibo/routes"
	"ginweibo/routes/named"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
)

func setupGin(g *gin.Engine) {
	gin.SetMode(config.AppConfig.RunMode)
	g.Static("/"+config.AppConfig.StaticPath, config.AppConfig.StaticPath)
	g.StaticFile("/favicon.ico", config.AppConfig.StaticPath+"/favicon.ico")
	// 注册模板函数
	g.SetFuncMap(template.FuncMap{
		"Mix":    helpers.Mix,
		"Static": helpers.Static,
		// 获取命名路由的 path
		"Route":         named.G,
		"RelativeRoute": named.GR,
	})
	// 模板存储路径
	g.LoadHTMLGlob(config.AppConfig.ViewsPath + "/**/*")
}

// 打印命名路由
func printRoute() {
	// 只有非 release 时才可用该函数
	if config.AppConfig.RunMode == config.RunmodeRelease {
		return
	}
	named.PrintRoutes()
}

func main() {
	// 解析命令行参数
	pflag.Parse()
	config.InitConfig()
	g := gin.New()
	setupGin(g)
	db := database.InitDB()
	db.AutoMigrate(
		&userModel.User{},
		&passwordResetModel.PasswordReset{},
		&statusModel.Status{},
		&followerModel.Follower{},
	)
	defer db.Close()
	// router register
	routes.Register(g)
	printRoute()
	// 启动
	fmt.Printf("\n\n----- Start to listening the incoming requests on http address: %s -----\n\n", config.AppConfig.Port)
	if err := http.ListenAndServe(config.AppConfig.Port, g); err != nil {
		log.Fatal("http server 启动失败", err)
	}
}
