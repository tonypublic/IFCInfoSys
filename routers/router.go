package routers

import (
	"github.com/astaxie/beego"
	"ifcinfosys/controllers"
)

func init() {
	//静态页面
	beego.Router("/", &controllers.MainController{})

	//api
	beego.Include(&controllers.CMSController{})
}
