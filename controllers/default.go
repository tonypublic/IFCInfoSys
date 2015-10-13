package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	records := new([10]recordBasicInfo)
	for i, _ := range records {
		records[i].Title = "走进股市"
		records[i].Date = "2015.4.14"
		records[i].Category = "微信 线上活动"
		records[i].Summary = "大风起兮云飞扬，热点轮换怎么玩，业内人士带你走进A股H股B股。"
		records[i].ItempicPath = "db/onlineactive/l00001/item.jpg"
	}

	c.Data["records"] = records

	c.TplNames = "index.html"
}

type recordBasicInfo struct {
	Title       string
	Date        string
	Category    string
	Summary     string
	ItempicPath string
}
