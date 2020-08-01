package home

import (
	// "fmt"
	_ "monteblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.get_article_class()
}

// func (c *BaseController) Finish() {...}

func (c *BaseController) get_article_class() {
	o := orm.NewOrm()
	var lists []orm.Params
	_, err := o.QueryTable("article_class").Values(&lists)
	if err == nil {
		// 查询每个分类中文章数量
		for k, v := range lists {
			tagCount, _ := o.QueryTable("article").Filter("article_class_id", v["ArticleclassId"]).Count()
			lists[k]["tagCount"] = tagCount
		}
		c.Data["articleClassList"] = lists
	}
}
