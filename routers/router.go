package routers

import (
	"note/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.SetStaticPath("images", "static/img")
	beego.SetStaticPath("css", "static/css/")
	beego.SetStaticPath("js", "static/js")
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.BlogController{})
}
