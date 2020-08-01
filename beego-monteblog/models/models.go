package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	ArticleId      int    `orm:"column(article_id);pk"`
	ArticleClassId int    `orm:"column(article_class_id)"`
	Title          string `orm:"size(255);column(title)"`
	Content        string `orm:"size(255);column(content);type(text)"`
	Time           int64  `orm:"column(time)"`
	UpdateTime     int64  `orm:"column(update_time)"`
}

type ArticleClass struct {
	ArticleclassId int    `orm:"column(article_class_id);pk"`
	Name           string `orm:"size(255);column(name)"`
}

type User struct {
	UserId   int    `orm:"column(user_id);pk"`
	Username string `orm:"size(32);column(username)"`
	Password string `orm:"size(32);column(password)"`
	Admin    int    `orm:"column(admin)"`
}

func init() {
	var user string = beego.AppConfig.String("mysqluser")
	var pass string = beego.AppConfig.String("mysqlpass")
	var urls string = beego.AppConfig.String("mysqlurls")
	var db string = beego.AppConfig.String("mysqldb")
	var databaseStr string = user + ":" + pass + "@tcp(" + urls + ")/" + db + "?charset=utf8"
	// 设置数据库链接
	orm.RegisterDataBase("default", "mysql", databaseStr, 30, 30)
	// 注册model
	orm.RegisterModel(new(ArticleClass), new(Article), new(User))
}
