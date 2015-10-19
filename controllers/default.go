package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"ifcinfosys/models"
	"strconv"
)

//活动列表
type recordBasicInfo struct {
	Title       string
	Date        string
	Category    string
	Summary     string
	ItempicPath string
}

//活动详细内容
type onLineActiveDetails struct {
	RecordNo string
	Category string
	Title    string
	Date     string
	Form     string
	Host     string
	Guest    string
	Summary  string
	Content  string
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	records := new([10]recordBasicInfo)
	for i, _ := range records {
		records[i].Title = "走进股市"
		records[i].Date = "2015.4.14"
		records[i].Category = "微信 | 线上活动"
		records[i].Summary = "大风起兮云飞扬，热点轮换怎么玩，业内人士带你走进A股H股B股。"
		records[i].ItempicPath = "db/onlineactive/l00001/thumbnails.jpg"
	}

	this.Data["records"] = records

	this.TplNames = "index.html"
}

//CMS API
type CMSController struct {
	beego.Controller
}

func (this *CMSController) URLMapping() {
	this.Mapping("api/list/", this.ListAddedRecords)
	this.Mapping("/api/details/", this.ViewDetails)
	//CRUD操作
	this.Mapping("/api/manage/add/", this.AddItem)
	this.Mapping("/api/manage/update", this.UpdateItem)
	this.Mapping("/api/manage/delete", this.DeleteItem)
}

// @router /api/list/:lastindex [get]
func (this *CMSController) ListAddedRecords() {
	lastindex, _ := this.GetInt(":lastindex")

	records := new([10]recordBasicInfo)
	for i, _ := range records {
		records[i].Title = "走进股市" + strconv.Itoa(i+lastindex)
		records[i].Date = "2015.4.14"
		records[i].Category = "微信 | 线上活动"
		records[i].Summary = "大风起兮云飞扬，热点轮换怎么玩，业内人士带你走进A股H股B股。"
		records[i].ItempicPath = "db/onlineactive/l00001/thumbnails.jpg"
	}

	this.Data["json"] = &records
	this.ServeJson()
}

// @router /api/details/:recordno [get]
func (this *CMSController) ViewDetails() {
	recordno, _ := this.GetInt(":recordno")
	fmt.Println(recordno)

	category := "线上活动"
	title := "走进股市"
	date := "2015.10.16"
	form := "微信"
	host := "张三，李四"
	guest := "张龙，赵虎"
	summary := "大风起兮云飞扬，热点轮换怎么玩，业内人士带你走进A股H股B股。"
	// picPath := "db/onlineactive/l00001/"

	content := "<p>Here is About page!</p>" +
		"<p>Go back or click  create-page to create dynamic page.</p>" +
		"<p>Mauris posuere sit amet metus id venenatis. Ut ante dolor, tempor nec commodo rutrum, varius at sem. Nullam ac nisi non neque ornare pretium. Nulla mauris mauris, consequat et elementum sit amet, egestas sed orci. In hac habitasse platea dictumst.</p>" +
		"<p align=\"center\"><img src=\"db/onlineactive/l00001/001.jpg\" width=\"100%\" height=\"auto\"></p>" +
		"<p>Fusce eros lectus, accumsan eget mi vel, iaculis tincidunt felis. Nulla tincidunt pharetra sagittis. Fusce in felis eros. Nulla sit amet aliquam lorem, et gravida ipsum. Mauris consectetur nisl non sollicitudin tristique. Praesent vitae metus ac quam rhoncus mattis vel et nisi. Aenean aliquet, felis quis dignissim iaculis, lectus quam tincidunt ligula, et venenatis turpis risus sed lorem. Morbi eu metus elit. Ut vel diam dolor."

	this.Data["title"] = title
	this.Data["date"] = date
	this.Data["category"] = form + " | " + category
	this.Data["content"] = "<p><strong>" + "嘉宾: " + host + "</strong></p>" +
		"<p><strong>" + "主持: " + guest + "</strong></p>" +
		"<p><i>" + summary + "</i></p>" +
		"<p>" + content + "</p></div>"

	this.TplNames = "details.html"
}

// @router /api/manage/add/:itemno [get]
func (this *CMSController) AddItem() {
	itemno := this.GetString(":itemno")
	fmt.Println(itemno)

	var item = new(models.OnLineActive)
	item.GetItem(itemno)

	this.Ctx.WriteString(itemno)
}

// @router /api/manage/update/:itemno [post]
func (this *CMSController) UpdateItem() {
	itemno := this.GetString(":itemno")
	fmt.Println(itemno)

}

// @router /api/manage/delete/:itemno [get]
func (this *CMSController) DeleteItem() {
	itemno := this.GetString(":itemno")
	fmt.Println(itemno)

}
