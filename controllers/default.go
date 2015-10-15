package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type recordBasicInfo struct {
	Title       string
	Date        string
	Category    string
	Summary     string
	ItempicPath string
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	records := new([10]recordBasicInfo)
	for i, _ := range records {
		records[i].Title = "走进股市"
		records[i].Date = "2015.4.14"
		records[i].Category = "微信 线上活动"
		records[i].Summary = "大风起兮云飞扬，热点轮换怎么玩，业内人士带你走进A股H股B股。"
		records[i].ItempicPath = "db/onlineactive/l00001/item.jpg"
	}

	this.Data["records"] = records

	this.TplNames = "index.html"
}

//CMS API
type CMSController struct {
	beego.Controller
}

func (this *CMSController) URLMapping() {
	this.Mapping("api/list", this.ListAddedRecords)
	this.Mapping("/api/details", this.ViewDetails)
}

// @router /api/list/:lastindex [get]
func (this *CMSController) ListAddedRecords() {
	lastindex, _ := this.GetInt(":lastindex")

	records := new([10]recordBasicInfo)
	for i, _ := range records {
		records[i].Title = "走进股市" + strconv.Itoa(i+lastindex)
		records[i].Date = "2015.4.14"
		records[i].Category = "微信 线上活动"
		records[i].Summary = "大风起兮云飞扬，热点轮换怎么玩，业内人士带你走进A股H股B股。"
		records[i].ItempicPath = "db/onlineactive/l00001/item.jpg"
	}

	this.Data["json"] = &records
	this.ServeJson()
}

// @router /api/details/:recordno [get]
func (this *CMSController) ViewDetails() {
	recordno, _ := this.GetInt(":recordno")
	fmt.Println(recordno)

	this.Data["content"] = "Here is About page!</p>" +
		"<p>Go back or click  create-page to create dynamic page.</p>" +
		"<p>Mauris posuere sit amet metus id venenatis. Ut ante dolor, tempor nec commodo rutrum, varius at sem. Nullam ac nisi non neque ornare pretium. Nulla mauris mauris, consequat et elementum sit amet, egestas sed orci. In hac habitasse platea dictumst.</p>" +
		"<p>Fusce eros lectus, accumsan eget mi vel, iaculis tincidunt felis. Nulla tincidunt pharetra sagittis. Fusce in felis eros. Nulla sit amet aliquam lorem, et gravida ipsum. Mauris consectetur nisl non sollicitudin tristique. Praesent vitae metus ac quam rhoncus mattis vel et nisi. Aenean aliquet, felis quis dignissim iaculis, lectus quam tincidunt ligula, et venenatis turpis risus sed lorem. Morbi eu metus elit. Ut vel diam dolor."

	this.TplNames = "details.html"
}
