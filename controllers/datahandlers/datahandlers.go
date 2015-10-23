/*
该包用来将数据库中取出的数据转化为前端需要展示的数据
*/
package datahandlers

import (
	"ifcinfosys/models"
	"strings"
)

//活动列表
type viewItem struct {
	No         string
	Title      string
	Date       string
	Category   string
	Summary    string
	Thumbnails string //信息缩略图路径
}

//活动详细内容
type ViewItemDetails struct {
	No       string
	Title    string
	Date     string
	Category string
	Content  string
	Head     string //文章头图片路径
}

//获取信息列表
func GetViewItems() (viewItems []viewItem, err error) {
	items, err := models.GetItems()
	if err == nil {
		vitem := new(viewItem)
		for _, item := range items {
			vitem.No = item.No
			vitem.Title = item.Title
			vitem.Date = strings.Fields(item.Updatetime)[0]
			vitem.Category = item.Type + " | " + item.Category
			vitem.Summary = item.Summary
			vitem.Thumbnails = getPathByType(item.Type) + item.No + "/thumbnails.jpg"

			viewItems = append(viewItems, *vitem)
		}
	}
	return
}

//根据记录编号获取记录的详细内容
func (this *ViewItemDetails) Get() (err error) {
	switch getItemType(this.No) {
	case "线上活动":
		err = getOnLineActiveDetail(this)
	case "项目信息":

	case "公告信息":

	default:
	}
	return
}

func getOnLineActiveDetail(vid *ViewItemDetails) (err error) {
	oad := &models.OnLineActive{No: vid.No}
	err = oad.Get()
	if err == nil {
		vid.Title = oad.Title
		vid.Date = strings.Fields(oad.Date)[0]
		vid.Category = getItemType(vid.No) + " | " + oad.Category
		vid.Content = "<p><strong>" + "嘉宾: " + oad.Hosts + "</strong></p>" +
			"<p><strong>" + "主持: " + oad.Guests + "</strong></p>" +
			"<p><i>" + oad.Summary + "</i></p>" +
			"<p>" + oad.Content + "</p></div>"
		vid.Head = getPathByType(getItemType(vid.No)) + vid.No + "/head.jpg"
	}
	return
}

//通过编号区别信息分类
func getItemType(no string) (t string) {
	switch no[0] {
	case 'l':
		t = "线上活动"
	case 'p':
		t = "项目信息"
	case 'b':
		t = "公告信息"
	default:
		t = "其他"
	}
	return
}

//通过信息分类获附件取存放目录
func getPathByType(t string) (dir string) {
	switch t {
	case "线上活动":
		dir = "db/onlineactive/"
	case "项目信息":
		dir = "db/project/"
	case "公告信息":
		dir = "db/bulletin/"
	}
	return
}
