package admin

import (
	"fmt"
	"monteblog/models"
	"monteblog/util"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type ArticleClassController struct {
	BaseController
}

func (c *ArticleClassController) List() {
	o := orm.NewOrm()
	var lists []orm.Params
	_, err := o.QueryTable("article_class").Values(&lists)
	if err == nil {
		c.Data["articleClassList"] = lists
	}
	c.TplName = "admin/article_class_list.html"
}

func (c *ArticleClassController) Show() {
	c.TplName = "admin/article_class_show.html"
}

func (c *ArticleClassController) Add() {
	id := c.Ctx.Input.Param(":id")
	o := orm.NewOrm()
	var article_class models.ArticleClass
	err := o.QueryTable("article_class").Filter("article_class_id", id).One(&article_class)
	if err == nil {
		c.Data["articleClass"] = article_class
	}
	c.TplName = "admin/article_class_add.html"
}

func (c *ArticleClassController) Del() {
	id, _ := c.GetInt(":id")
	fmt.Print(id, "\n")
	o := orm.NewOrm()
	if _, err := o.Delete(&models.ArticleClass{ArticleclassId: id}); err == nil {
		c.Redirect("/admin/article-list", 302)
	}
}

func (c *ArticleClassController) AjaxUpdate() {
	o := orm.NewOrm()
	name := c.GetString("name")
	if name == "" {
		data := util.Error{Code: 0, Msg: "名称不能为空", Data: ""}
		c.Data["json"] = data
		c.ServeJSON()
	}
	id := c.GetString("articleClassId")
	if id == "" {
		articleclassAdd := new(models.ArticleClass)
		articleclassAdd.Name = name
		_, err := o.Insert(articleclassAdd)
		if err != nil {
			data := util.Error{Code: 0, Msg: "数据添加失败", Data: err}
			c.Data["json"] = data
			c.ServeJSON()
		}
		data := util.Normal{Code: 1, Msg: "数据添加成功", Data: ""}
		c.Data["json"] = data
		c.ServeJSON()
	} else {
		intid, _ := strconv.Atoi(id)
		articleclassUp := models.ArticleClass{ArticleclassId: intid}
		if o.Read(&articleclassUp) == nil {
			articleclassUp.Name = name
			if _, err := o.Update(&articleclassUp); err == nil {
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
