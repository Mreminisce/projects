package main

import (
	_ "monteblog/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 加载data.conf
	// beego.LoadAppConfig("ini", "conf/data.conf")
	beego.Run()
}
