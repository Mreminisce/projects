package config

import "github.com/spf13/viper"

type mailConfig struct {
	Driver   string // smtp、log(log 把邮件写在日志中，便于调试)
	Host     string // 邮箱的服务器地址
	Port     int    // 邮箱的服务器端口
	User     string // 发送者的 name
	Password string // 授权码或密码
	FromName string // 邮件发送者名称
}

func newMailConfig() *mailConfig {
	viper.SetDefault("MAIL.MAIL_DRIVER", "ginweibo")
	viper.SetDefault("MAIL.MAIL_HOST", "")
	viper.SetDefault("MAIL.MAIL_PORT", 25)
	viper.SetDefault("MAIL.MAIL_USERNAME", "")
	viper.SetDefault("MAIL.MAIL_PASSWORD", "")
	viper.SetDefault("MAIL.MAIL_FROM_NAME", "ginweibo")
	return &mailConfig{
		Driver:   viper.GetString("MAIL.MAIL_DRIVER"),
		Host:     viper.GetString("MAIL.MAIL_HOST"),
		Port:     viper.GetInt("MAIL.MAIL_PORT"),
		User:     viper.GetString("MAIL.MAIL_USERNAME"),
		Password: viper.GetString("MAIL.MAIL_PASSWORD"),
		FromName: viper.GetString("MAIL.MAIL_FROM_NAME"),
	}
}
