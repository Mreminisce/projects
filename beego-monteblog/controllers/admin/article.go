package admin

import (
	"monteblog/models"
	"monteblog/util"
	"time"

	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	BaseController
}

func (c *ArticleController) List() {
	o := orm.NewOrm()
	var lists []orm.Params
	_, err := o.QueryTable("article").Values(&lists)
	if err == nil {
		var article models.ArticleClass
		for key, value := range lists {
			o.QueryTable("article_class").Filter("article_class_id", value["ArticleClassId"]).One(&article, "Name")
			lists[key]["ArticleClassName"] = article.Name
		}
		c.Data["articleList"] = lists
	}
	c.TplName = "admin/article_list.html"
}

func (c *ArticleController) Add() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	var article models.Article
	err := o.QueryTable("article").Filter("article_id", id).One(&article)
	if err == nil {
		c.Data["article"] = article
	}
	var articlelists []orm.Params
	_, articleerr := o.QueryTable("article_class").Values(&articlelists)
	if articleerr == nil {
		for key, value := range articlelists {
			if int64(article.ArticleClassId) == value["ArticleclassId"] {
				articlelists[key]["ArticleclassType"] = 1
			} else {
				articlelists[key]["ArticleclassType"] = 0
			}
		}
		c.Data["articleClassList"] = articlelists
	}
	c.TplName = "admin/article_add.html"
}

func (c *ArticleController) Del() {
	id, _ := c.GetInt(":id")
	o := orm.NewOrm()
	if _, err := o.Delete(&models.Article{ArticleId: id}); err == nil {
		c.Redirect("/admin/articlelist", 302)
	}
}

func (c *ArticleController) AjaxUpdate() {
	title := c.GetString("title")
	if title == "" {
		c.Data["json"] = util.Error{Code: 0, Msg: "文章标题不能为空", Data: ""}
		c.ServeJSON()
	}
	article_class_id, article_class_err := c.GetInt("article_class_id")
	if article_class_err != nil {
		c.Data["json"] = util.Error{Code: 0, Msg: "文章分类不能为空", Data: ""}
		c.ServeJSON()
	}
	content := c.GetString("content")
	if content == "" {
		c.Data["json"] = util.Error{Code: 0, Msg: "内容不能为空", Data: ""}
		c.ServeJSON()
	}
	article_id, err := c.GetInt("article_id")
	o := orm.NewOrm()
	if err != nil {
		articleAdd := new(models.Article)
		articleAdd.Title = title
		articleAdd.ArticleClassId = article_class_id
		articleAdd.Content = content
		articleAdd.Time = time.Now().Unix()
		_, err := o.Insert(articleAdd)
		if err != nil {
			data := util.Error{Code: 0, Msg: "数据添加失败", Data: err}
			c.Data["json"] = data
			c.ServeJSON()
		}
		data := util.Normal{Code: 1, Msg: "数据添加成功", Data: ""}
		c.Data["json"] = data
		c.ServeJSON()
	} else {
		articleUp := models.Article{ArticleId: article_id}
		if o.Read(&articleUp) == nil {
			articleUp.Title = title
			articleUp.ArticleClassId = article_class_id
			articleUp.Content = content
			articleUp.UpdateTime = time.Now().Unix()
			if _, err := o.Update(&articleUp); err == nil {
				data := util.Normal{Code: 1, Msg: "数据编辑成功", Data: ""}
				c.Data["json"] = data
			} else {
				data := util.Normal{Code: 0, Msg: "数据编辑失败", Data: ""}
				c.Data["json"] = data
			}
		} else {
			data := util.Normal{Code: 0, Msg: "要编辑的数据不存在", Data: ""}
			c.Data["json"] = data
		}
		c.ServeJSON()
	}
}
