package config

import (
	"github.com/spf13/viper"
)

type appConfig struct {
	Name                      string
	RunMode                   string // 运行模式: debug, release, test
	Port                      string // 运行端口
	URL                       string // 完整 url
	Key                       string // secret key
	StaticPath                string // 静态资源存放路径
	ResourcesPath             string // 模板等前端源码文件存放路径
	ViewsPath                 string // 模板文件存放的路径
	EnableCsrf                bool
	CsrfParamName             string
	CsrfHeaderName            string
	AuthSessionKey            string // session key
	ContextCurrentUserDataKey string // Context 中当前用户数据的 key
}

func newAppConfig() *appConfig {
	viper.SetDefault("APP.NAME", "ginweibo")
	viper.SetDefault("APP.RUNMODE", "release")
	viper.SetDefault("APP.PORT", ":8088")
	viper.SetDefault("APP.URL", "")
	viper.SetDefault("APP.KEY", "base64:O+VQ74YEigLPDzLKnh2HW/yjCdU2ON9v7xuKBgSOEAo=")
	viper.SetDefault("APP.ENABLE_CSRF", true)
	return &appConfig{
		Name:                      viper.GetString("APP.NAME"),
		RunMode:                   viper.GetString("APP.RUNMODE"),
		Port:                      viper.GetString("APP.PORT"),
		URL:                       viper.GetString("APP.URL"),
		Key:                       viper.GetString("APP.KEY"),
		StaticPath:                "static",
		ResourcesPath:             "resources",
		ViewsPath:                 "resources/views",
		EnableCsrf:                viper.GetBool("APP.ENABLE_CSRF"),
		CsrfParamName:             "_csrf",
		CsrfHeaderName:            "X-CsrfToken",
		AuthSessionKey:            "gin_session",
		ContextCurrentUserDataKey: "currentUserData",
	}
}
