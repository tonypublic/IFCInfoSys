package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"ifcinfosys/controllers/datahandlers"
	// "strconv"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	items, err := datahandlers.GetViewItems()
	if err == nil {
		this.Data["items"] = items
		this.TplNames = "index.html"
	}
}

//CMS API
type CMSController struct {
	beego.Controller
}

func (this *CMSController) URLMapping() {
	this.Mapping("/api/list/", this.ListAddedRecords)
	this.Mapping("/api/details/", this.ViewDetails)
	//CRUD操作
	this.Mapping("/api/manage/add/", this.AddItem)
	this.Mapping("/api/manage/update", this.UpdateItem)
	this.Mapping("/api/manage/delete", this.DeleteItem)
}

// @router /api/list/:lastindex [get]
func (this *CMSController) ListAddedRecords() {
	// lastindex, _ := this.GetInt(":lastindex")

	items, err := datahandlers.GetViewItems()
	if err == nil {
		this.Data["json"] = &items
		this.ServeJson()
	}
}

// @router /api/details/:itemno [get]
func (this *CMSController) ViewDetails() {
	itemno := this.GetString(":itemno")
	itemDetails := &datahandlers.ViewItemDetails{No: itemno}
	err := itemDetails.Get()
	if err == nil {
		this.Data["itemdetails"] = itemDetails
		this.TplNames = "details.html"
	}
}

// @router /api/manage/add/:itemno [get]
func (this *CMSController) AddItem() {

}

// @router /api/manage/update/:itemno [post]
func (this *CMSController) UpdateItem() {

}

// @router /api/manage/delete/:itemno [get]
func (this *CMSController) DeleteItem() {

}
