package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type SummaryList struct {
	No         string
	Type       string
	Updatetime string
	Pageviewed int
}

type OnLineActive struct {
	No       string
	Title    string
	Category string
	Date     string
	Hosts    string
	Guests   string
	Summary  string
	Content  string
}

type Bulletin struct {
	No        string
	Title     string
	Category  string
	Begintime string
	Endtime   string
	Contacts  string
	Summary   string
	Content   string
}

type Project struct {
	No        string
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
	orm.RegisterModel(new(SummaryList), new(OnLineActive), new(Bulletin), new(Project))
}
