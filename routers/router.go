package routers

import (
	"education/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

func init() {

	// 默认首页主体
    beego.Router("/", &controllers.AdminController{})

	// Admin自动匹配  例如  /admin/login
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.AuthController{})

	// Admin 登陆验证
	beego.InsertFilter("/admin/*", beego.BeforeRouter, adminLoginFilter)
}

func adminLoginFilter(ctx *context.Context)  {

	URI := ctx.Request.RequestURI

	if strings.Index(URI, "/admin/login") == 0{
		return
	}
}