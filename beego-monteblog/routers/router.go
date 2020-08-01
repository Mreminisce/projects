package routers

import (
	"monteblog/controllers/admin"
	"monteblog/controllers/home"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &home.IndexController{}, "get:Index")
	beego.Router("/about", &home.AboutController{}, "get:Show")
	beego.Router("/article/list/:cid:int/?:page:int", &home.ArticleController{}, "get:List")
	beego.Router("/article/show/:id:int", &home.ArticleController{}, "get:Show")
	beego.Router("/article/search/:keyword/?:page:int", &home.ArticleController{}, "get:Search")
	beego.Router("/admin/index", &admin.IndexController{}, "get:Index")
	beego.Router("/admin/login", &admin.LoginController{}, "get,post:Login")
	beego.Router("/admin/signout", &admin.LoginController{}, "get:Signout")
	beego.Router("/admin/articlelist", &admin.ArticleController{}, "get:List")
	beego.Router("/admin/articledel/:id:int", &admin.ArticleController{}, "get:Del")
	beego.Router("/admin/articleadd/?:id:int", &admin.ArticleController{}, "get:Add")
	beego.Router("/admin/articleupdate", &admin.ArticleController{}, "post:AjaxUpdate")
	beego.Router("/admin/article-list", &admin.ArticleClassController{}, "get:List")
	beego.Router("/admin/article-del/:id:int", &admin.ArticleClassController{}, "get:Del")
	beego.Router("/admin/article-add/?:id:int", &admin.ArticleClassController{}, "get:Add")
	beego.Router("/admin/article-update", &admin.ArticleClassController{}, "post:AjaxUpdate")
	// 富文本编辑器上传图片接口
	beego.Router("/go/upload_json", &admin.UploadController{}, "post:UploadJson")
}
