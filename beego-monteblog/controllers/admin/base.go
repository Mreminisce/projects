package admin

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Data["custom_xsrf"] = c.XSRFToken()
	c.HasSession()
}

func (c *BaseController) HasSession() {
	userinfo := c.GetSession("userinfo")
	if userinfo == nil {
		c.Redirect("/admin/login", 302)
		return
	}
}
