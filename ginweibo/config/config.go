package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

const (
	RunmodeDebug   = "debug"
	RunmodeRelease = "release"
	RunmodeTest    = "test"
	configFilePath = "./config.yaml"
	logFilePath    = "log/ginweibo.log"
	configFileType = "yaml"
)

var (
	AppConfig  *appConfig
	DBConfig   *dbConfig
	MailConfig *mailConfig
)

// 监控配置文件变化
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(ev fsnotify.Event) {
		log.Infof("Config file changed: %s", ev.Name)
	})
}

func InitConfig() {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败，请检查: %v", err))
	}
	initLog()
	AppConfig = newAppConfig()
	DBConfig = newDBConfig()
	MailConfig = newMailConfig()
	// 热更新配置文件
	watchConfig()
}
