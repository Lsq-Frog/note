package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["note/controllers:BlogController"] = append(beego.GlobalControllerRouter["note/controllers:BlogController"],
		beego.ControllerComments{
			Method: "About",
			Router: `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["note/controllers:BlogController"] = append(beego.GlobalControllerRouter["note/controllers:BlogController"],
		beego.ControllerComments{
			Method: "Blog",
			Router: `/blog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["note/controllers:BlogController"] = append(beego.GlobalControllerRouter["note/controllers:BlogController"],
		beego.ControllerComments{
			Method: "Contact",
			Router: `/contact`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["note/controllers:BlogController"] = append(beego.GlobalControllerRouter["note/controllers:BlogController"],
		beego.ControllerComments{
			Method: "Home",
			Router: `/home`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["note/controllers:BlogController"] = append(beego.GlobalControllerRouter["note/controllers:BlogController"],
		beego.ControllerComments{
			Method: "Support",
			Router: `/support`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
