package home

import (
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Index() {
	o := orm.NewOrm()
	// 查询文章列表
	var article_list []orm.Params
	_, article_err := o.QueryTable("article").OrderBy("-article_id").Values(&article_list)
	if article_err == nil {
		c.Data["articleList"] = article_list
	}
	c.TplName = "home/index.html"
}
