package system

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	DSN           string `yaml:"dsn"`            //数据库地址
	PageSize      int    `yaml:"page_size"`      //尺寸大小
	Port          string `yaml:"addr"`           //请求端口
	SessionSecret string `yaml:"session_secret"` //session
	SignupEnabled bool   `yaml:"signup_enabled"` // 是否启动注册
	Public        string `yaml:"public"`         //公开
	Domain        string `yaml:"domain"`         //域名
}

const (
	DEFAULT_PAGESIZE = 10
)

var configuration *Configuration

func GetConfiguration() *Configuration {
	return configuration
}

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	// 获取文件中的所有数据
	var config Configuration
	// 配置文件解析
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil
	}
	if config.PageSize <= 0 {
		config.PageSize = DEFAULT_PAGESIZE
	}
	configuration = &config
	return err
}
