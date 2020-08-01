package home

type AboutController struct {
	BaseController
}

// get 请求 about/show
func (c *AboutController) Show() {
	c.TplName = "home/about.html"
}
