package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["note/controllers:BlogController"] = append(beego.GlobalControllerRouter["note/controllers:BlogController"],
		beego.ControllerComments{
			"Home",
			`/home`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["note/controllers:BlogController"] = append(beego.GlobalControllerRouter["note/controllers:BlogController"],
		beego.ControllerComments{
			"Support",
			`/support`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["note/controllers:BlogController"] = append(beego.GlobalControllerRouter["note/controllers:BlogController"],
		beego.ControllerComments{
			"About",
			`/about`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["note/controllers:BlogController"] = append(beego.GlobalControllerRouter["note/controllers:BlogController"],
		beego.ControllerComments{
			"Blog",
			`/blog`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["note/controllers:BlogController"] = append(beego.GlobalControllerRouter["note/controllers:BlogController"],
		beego.ControllerComments{
			"Contact",
			`/contact`,
			[]string{"get"},
			nil})

}
