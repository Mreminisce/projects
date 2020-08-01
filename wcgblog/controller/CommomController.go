package controller

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"wcgblog/service"
	"wcgblog/system"
	"wcgblog/utils"

	"github.com/denisbakhtin/sitemap"
	"github.com/gin-gonic/gin"
)

const (
	SESSION_KEY      = "UserID"      // session key
	CONTEXT_USER_KEY = "User"        // 上下文连接
	SESSION_CAPTCHA  = "GIN_CAPTCHA" // 验证码密匙
)

func Handle404(c *gin.Context) {
	HandleMessage(c, "Sorry,I lost myself!")
}
func HandleMessage(c *gin.Context, message string) {
	c.HTML(http.StatusNotFound, "errors/error.html", gin.H{
		"message": message,
	})
}

// 每天获取一次
func CreateXMLSitemap() {
	configuration := system.GetConfiguration()
	// path.join 将输入的拼接起来 如path.Join("c:", "aa", "bb", "cc.txt"))         //c:/aa/bb/cc.txt
	join := path.Join(configuration.Public, "sitemap")
	// os.mkdirall 创建多级目录
	os.MkdirAll(join, os.ModePerm)
	domin := configuration.Domain
	// 获取北京时间
	time := utils.GetCurrentTime()
	items := make([]sitemap.Item, 0)
	// 网页地图 SiteMap（站点地图） 是一个列出你网站网页的文件，来告知 Google 和其他搜索引擎您网站内容的组织情况。 Googlebot 等搜索引擎网络抓取工具读取此文件，以更智能地抓取您的网站。
	items = append(items, sitemap.Item{
		Loc:        domin,
		LastMod:    time,
		Changefreq: "daily",
		Priority:   1,
	})
	posts, err := service.ListPublishedPost("", 0, 0)
	if err == nil {
		for _, post := range posts {
			items = append(items, sitemap.Item{
				Loc:        fmt.Sprintf("%s/post/%d", domin, post.ID),
				LastMod:    post.UpdatedAt,
				Changefreq: "weekly",
				Priority:   0.9,
			})
		}
	}
	pages, err := service.ListPubilshePage()
	if err == nil {
		for _, page := range pages {
			items = append(items, sitemap.Item{
				Loc:        fmt.Sprintf("%s/page/%d", domin, page.ID),
				LastMod:    page.UpdatedAt,
				Changefreq: "monthly",
				Priority:   0.8,
			})
		}
	}
	if err := sitemap.SiteMap(path.Join(join, "sitemap.xml.ge"), items); err != nil {
		return
	}
	if err := sitemap.SiteMapIndex(join, "sitemap_index.xml", domin+"/static/sitemap/"); err != nil {
		return
	}
}
