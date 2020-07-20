package helpers

import (
	"encoding/json"
	"os"
	"path"

	"ginweibo/config"
)

// 存储 mix-manifest.json 解析出来的 path map
var manifests = make(map[string]string)

// 生成项目静态文件地址
func Static(staticFilePath string) string {
	return "/" + config.AppConfig.StaticPath + staticFilePath
}

// 根据 laravel-mix 的 static/mix-manifest.json 生成静态文件 path
func Mix(staticFilePath string) string {
	result := manifests[staticFilePath]
	if result == "" {
		filename := path.Join(config.AppConfig.StaticPath, "mix-manifest.json")
		file, err := os.Open(filename)
		if err != nil {
			return Static(staticFilePath)
		}
		defer file.Close()
		dec := json.NewDecoder(file)
		if err := dec.Decode(&manifests); err != nil {
			return Static(staticFilePath)
		}
		result = manifests[staticFilePath]
	}
	if result == "" {
		return Static(staticFilePath)
	}
	return Static(result)
}
