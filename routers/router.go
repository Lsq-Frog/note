package routers

import (
	"note/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.SetStaticPath("images", "static/img/")
	beego.SetStaticPath("css", "static/css/")
	beego.SetStaticPath("js", "static/js/")
	beego.SetStaticPath("fonts", "static/fonts/")
	beego.SetStaticPath("video", "static/video/")
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.BlogController{})
}
