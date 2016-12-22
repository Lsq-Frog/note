package routers

import (
	"note/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.SetStaticPath("images", "static/img")
	beego.SetStaticPath("css", "static/css/")
	beego.SetStaticPath("js", "static/js")
	beego.Include(&controllers.BlogController{})
}
