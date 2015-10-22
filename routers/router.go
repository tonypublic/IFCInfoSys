package routers

import (
	"github.com/astaxie/beego"
	"ifcinfosys/controllers"
)

func init() {
	//首页
	beego.Router("/", &controllers.MainController{})

	//api
	beego.Include(&controllers.CMSController{})
}
