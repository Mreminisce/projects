package admin

import (
	"html/template"
	"monteblog/models"
	"monteblog/util"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username != "" {
		o := orm.NewOrm()
		user := models.User{Username: username, Password: util.Md5V(password)}
		err := o.Read(&user, "Username", "Password")
		if err == orm.ErrNoRows {
			c.Redirect("/admin/login", 302)
		}
		// 写入用户数据到 session
		c.SetSession("userinfo", user)
		c.Redirect("/admin/index", 302)
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "admin/login.html"
}

func (c *LoginController) Signout() {
	c.DelSession("userinfo")
	c.Redirect("/admin/index", 302)
}
