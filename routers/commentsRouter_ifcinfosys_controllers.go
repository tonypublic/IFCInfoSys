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
			`/api/details/:itemno`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ifcinfosys/controllers:CMSController"] = append(beego.GlobalControllerRouter["ifcinfosys/controllers:CMSController"],
		beego.ControllerComments{
			"AddItem",
			`/api/manage/add/:itemno`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["ifcinfosys/controllers:CMSController"] = append(beego.GlobalControllerRouter["ifcinfosys/controllers:CMSController"],
		beego.ControllerComments{
			"UpdateItem",
			`/api/manage/update/:itemno`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["ifcinfosys/controllers:CMSController"] = append(beego.GlobalControllerRouter["ifcinfosys/controllers:CMSController"],
		beego.ControllerComments{
			"DeleteItem",
			`/api/manage/delete/:itemno`,
			[]string{"get"},
			nil})

}
