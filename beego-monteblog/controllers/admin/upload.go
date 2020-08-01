package admin

import (
	"log"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) UploadJson() {
	f, h, err := c.GetFile("imgFile")
	if err != nil {
		log.Fatal("getfile err ", err)
		jsonDataE := make(map[string]interface{})
		jsonDataE["error"] = 1
		jsonDataE["url"] = "上传失败"
		c.Data["json"] = jsonDataE
		c.ServeJSON()
	}
	defer f.Close()
	// 保存位置在 static/upload, 没有文件夹要先创建
	c.SaveToFile("imgFile", "static/upload/"+h.Filename)
	jsonDataO := make(map[string]interface{})
	jsonDataO["error"] = 0
	jsonDataO["url"] = "/static/upload/" + h.Filename
	c.Data["json"] = jsonDataO
	c.ServeJSON()
}

func (c *UploadController) FileUploadJson() {
	c.Data["json"] = "111.jpg"
	c.ServeJSON()
}
