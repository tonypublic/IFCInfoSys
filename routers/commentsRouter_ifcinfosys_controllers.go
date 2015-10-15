package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["ifcinfosys/controllers:CMSController"] = append(beego.GlobalControllerRouter["ifcinfosys/controllers:CMSController"],
		beego.ControllerComments{
			"ListAddedRecords",
			`/api/list/:lastindex`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ifcinfosys/controllers:CMSController"] = append(beego.GlobalControllerRouter["ifcinfosys/controllers:CMSController"],
		beego.ControllerComments{
			"ViewDetails",
			`/api/details/:recordno`,
			[]string{"get"},
			nil})

}
