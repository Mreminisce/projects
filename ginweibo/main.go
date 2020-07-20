package main

import (
	"fmt"
	"ginweibo/config"
	"ginweibo/database"
	followerModel "ginweibo/models/follower"
	passwordResetModel "ginweibo/models/password_reset"
	statusModel "ginweibo/models/status"
	userModel "ginweibo/models/user"
	"ginweibo/pkg/helpers"
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
	g.SetFuncMap(template.FuncMap{
		"Mix":           helpers.Mix,
		"Static":        helpers.Static,
		"Route":         named.G, // 获取命名路由的 path
		"RelativeRoute": named.GR,
	})
	g.LoadHTMLGlob(config.AppConfig.ViewsPath + "/**/*")
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
	routes.Register(g)
	fmt.Printf("\n\n----- Start to listening the incoming requests on http address: %s -----\n\n", config.AppConfig.Port)
	if err := http.ListenAndServe(config.AppConfig.Port, g); err != nil {
		log.Fatal("http server 启动失败", err)
	}
}
