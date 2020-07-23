package controllers

import (
	"fmt"
	"ginweibo/config"
	"ginweibo/middleware/flash"
	"ginweibo/utils/view"
	"ginweibo/routes/named"
	"html/template"
	"math"
	"net/http"
	"strconv"

	"ginweibo/middleware/auth"
	viewmodels "ginweibo/middleware/viewmodels"

	"github.com/gin-gonic/gin"
)

type renderObj = map[string]interface{}

const (
	csrfInputHTML = "csrfField"
	csrfTokenName = "csrfToken"
)

func csrfField(c *gin.Context) (template.HTML, string, bool) {
	token := c.Keys[config.AppConfig.CsrfParamName]
	tokenStr, ok := token.(string)
	if !ok {
		return "", "", false
	}
	return template.HTML(fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`, config.AppConfig.CsrfParamName, tokenStr)), tokenStr, true
}

// Render : 渲染 html
func Render(c *gin.Context, tplPath string, data renderObj) {
	obj := make(renderObj)
	flashStore := flash.Read(c)
	oldValueStore := flash.ReadOldFormValue(c)
	validateMsgArr := flash.ReadValidateMessage(c)
	// flash 数据
	obj[flash.FlashInContextAndCookieKeyName] = flashStore.Data
	// 上次 post form 的数据，用于回填
	obj[flash.OldValueInContextAndCookieKeyName] = oldValueStore.Data
	// 上次表单的验证信息
	obj[flash.ValidateContextAndCookieKeyName] = validateMsgArr
	if config.AppConfig.EnableCsrf {
		if csrfHTML, csrfToken, ok := csrfField(c); ok {
			obj[csrfInputHTML] = csrfHTML
			obj[csrfTokenName] = csrfToken
		}
	}
	// 获取当前登录的用户 (如果用户登录了的话，中间件中会通过 session 存储用户数据)
	if user, err := auth.GetCurrentUserFromContext(c); err == nil {
		obj[config.AppConfig.ContextCurrentUserDataKey] = viewmodels.NewUserViewModelSerializer(user)
	}
	// 填充传递进来的数据
	for k, v := range data {
		obj[k] = v
	}
	c.HTML(http.StatusOK, tplPath, obj)
}

// RenderError : 渲染错误页面
func RenderError(c *gin.Context, code int, msg string) {
	errorCode := code
	if code == 419 || code == 403 {
		errorCode = 403
	}
	c.HTML(code, "error/error.html", gin.H{
		"errorMsg":  msg,
		"errorCode": errorCode,
		"errorImg":  view.Static("/svg/" + strconv.Itoa(code) + ".svg"),
		"backUrl":   named.G("root"),
	})
}

func Render403(c *gin.Context, msg string) {
	RenderError(c, http.StatusForbidden, msg)
}

func Render404(c *gin.Context) {
	RenderError(c, http.StatusNotFound, "很抱歉！您浏览的页面不存在。")
}

func RenderUnauthorized(c *gin.Context) {
	Render403(c, "很抱歉，您没有权限访问该页面")
}

func redirect(c *gin.Context, redirectPath string) {
	// 千万注意，这个地方不能用 301(永久重定向)
	c.Redirect(http.StatusFound, redirectPath)
}

// Redirect : 路由重定向 use path
func Redirect(c *gin.Context, redirectPath string, withRoot bool) {
	path := redirectPath
	if withRoot {
		path = config.AppConfig.URL + redirectPath
	}
	redirect(c, path)
}

// RedirectRouter : 路由重定向 use router name
func RedirectRouter(c *gin.Context, routerName string, args ...interface{}) {
	redirect(c, named.G(routerName, args...))
}

// RedirectToLoginPage : 重定向到登录页面
func RedirectToLoginPage(c *gin.Context) {
	loginPath := named.G("login.create")
	if c.Request.Method == http.MethodPost {
		redirect(c, loginPath)
		return
	}
	redirect(c, loginPath+"?back="+c.Request.URL.Path)
}

// GetIntParam : 从 path params 中获取 int 参数
// http://a.com/xx/1 => 获取到 int 1
func GetIntParam(c *gin.Context, key string) (int, error) {
	i, err := strconv.Atoi(c.Param(key))
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetPageQuery 从 query 中获取有关分页的参数
// xx.com?page=1&pageline=10
func GetPageQuery(c *gin.Context, defaultPageLine, totalCount int) (offset, limit, currentPage, pageTotalCount int) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	currentPage = page
	pageline, err := strconv.Atoi(c.Query("pageline"))
	if err != nil {
		pageline = defaultPageLine
	}
	page = page - 1
	if page == 0 {
		offset = 0
	} else {
		offset = page * pageline
	}
	limit = pageline
	pageTotalCount = int(math.Ceil(float64(totalCount) / float64(pageline)))
	if pageTotalCount <= 0 {
		pageTotalCount = 1
	}
	return
}
