package models

import (
	"fmt"
	"github.com/astaxie/beego/cache"
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

//全局缓存
var bm cache.Cache

func init() {
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)

	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "sqlite3", "db/ifcdb.db", maxIdle, maxConn)

	// 需要在init中注册定义的model
	orm.RegisterModel(new(VItemList), new(OnLineActive))

	//设定一个缓存，保存最初始的ItemList，避免每次打开主页就去读取数据库
	bm, _ = cache.NewCache("memory", `{"interval":60}`)
}

//获取信息列表, pos为查询的起始记录号
func GetItems(pos int) (items []VItemList, err error) {
	//如果是首页，先查看缓存，缓存中不存在再从数据库中取数据
	if pos == 0 && bm.IsExist("vitemlists") {
		items, _ = bm.Get("vitemlists").([]VItemList)
	} else {
		qs := orm.NewOrm().QueryTable("v_item_list")

		//每次取出的记录数为7
		_, err = qs.Limit(7, pos).All(&items)
		bm.Put("vitemlists", items, 60)
	}
	return
}

//返回一项记录的详细信息
func (this *OnLineActive) Get() (err error) {
	o := orm.NewOrm()
	err = o.Read(this)
	return
}
