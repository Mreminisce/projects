package home

import (
	"monteblog/models"
	"time"

	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	BaseController
}

// 按分类查询文章
func (c *ArticleController) List() {
	cid := c.Ctx.Input.Param(":cid")
	// page := c.Ctx.Input.Param(":page")
	page, _ := c.GetInt(":page")
	// 查询分类名称 + id
	o := orm.NewOrm()
	var article_class models.ArticleClass
	err := o.QueryTable("article_class").Filter("article_class_id", cid).One(&article_class)
	if err == nil {
		c.Data["articleClass"] = article_class
	}
	// 查询文章列表
	var article_list []orm.Params
	_, article_err := o.QueryTable("article").Filter("article_class_id", cid).Limit(6, (page-1)*6).OrderBy("-article_id").Values(&article_list)
	if article_err == nil {
		c.Data["articleList"] = article_list
	}
	c.TplName = "home/article_list.html"
}

// 文章详情
func (c *ArticleController) Show() {
	id := c.Ctx.Input.Param(":id")
	o := orm.NewOrm()
	var article models.Article
	err := o.QueryTable("article").Filter("article_id", id).One(&article)
	if err == nil {
		c.Data["article"] = article
	}
	// 获取分类名
	var article_class models.ArticleClass
	article_class_err := o.QueryTable("article_class").Filter("article_class_id", article.ArticleClassId).One(&article_class)
	if article_class_err == nil {
		c.Data["articleClass"] = article_class
	}
	c.TplName = "home/article_show.html"
}

func (c *ArticleController) Search() {
	var starttime, endtime int64
	page, _ := c.GetInt(":page")
	keyword := c.GetString(":keyword")
	// 查询开始时间戳纳秒
	starttime = time.Now().UnixNano()
	o := orm.NewOrm()
	var ormQuery orm.QuerySeter
	ormQuery = o.QueryTable("article").Filter("title__icontains", keyword).Filter("content__icontains", keyword).OrderBy("-article_id")
	// 查询总条数
	article_count, _ := ormQuery.Count()
	// 查询分页后的数据
	var article_list []orm.Params
	_, article_err := ormQuery.Limit(6, (page-1)*6).Values(&article_list)
	if article_err == nil {
		c.Data["articleList"] = article_list
	}
	endtime = time.Now().UnixNano()
	totaltime := float64(endtime-starttime) / 1000000000
	c.Data["totaltime"] = totaltime        // 总用时
	c.Data["articleCount"] = article_count // 总条数
	c.Data["keyword"] = keyword
	c.Data["page"] = page
	c.TplName = "home/search.html"
}
