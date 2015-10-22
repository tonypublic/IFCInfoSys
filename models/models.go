package models

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type VItemList struct {
	No         string `orm:"pk"`
	Type       string
	Updatetime string
	Pageviewed int
	Title      string
	Category   string
	Summary    string
}

type OnLineActive struct {
	No       string `orm:"pk"`
	Title    string
	Category string
	Date     string
	Hosts    string
	Guests   string
	Summary  string
	Content  string
}

type Bulletin struct {
	No        string `orm:"pk"`
	Title     string
	Category  string
	Begintime string
	Endtime   string
	Contacts  string
	Summary   string
	Content   string
}

type Project struct {
	No        string `orm:"pk"`
	Title     string
	Category  string
	Begintime string
	Endtime   string
	Contacts  string
	Summary   string
	Content   string
	Phase     string
}

func init() {
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)

	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "sqlite3", "db/ifcdb.db", maxIdle, maxConn)

	// 需要在init中注册定义的model
	orm.RegisterModel(new(VItemList), new(OnLineActive))
}

//获取信息列表
func GetItems() (items []VItemList, err error) {
	qs := orm.NewOrm().QueryTable("v_item_list")
	_, err = qs.Limit(10, 0).All(&items)
	return
}

//返回一项记录的详细信息
func (this *OnLineActive) Get() (err error) {
	o := orm.NewOrm()
	err = o.Read(this)
	return
}
