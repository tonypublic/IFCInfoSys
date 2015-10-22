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
	//api
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

	// records := new([10]recordBasicInfo)
	// for i, _ := range records {
	// 	records[i].Title = "走进股市" + strconv.Itoa(i+lastindex)
	// 	records[i].Date = "2015.4.14"
	// 	records[i].Category = "微信 | 线上活动"
	// 	records[i].Summary = "大风起兮云飞扬，热点轮换怎么玩，业内人士带你走进A股H股B股。"
	// 	records[i].ItempicPath = "db/onlineactive/l00001/thumbnails.jpg"
	// }

	// this.Data["json"] = &records
	// this.ServeJson()
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
