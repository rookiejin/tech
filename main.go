package main

import (
	_ "tech/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"tech/models"
)

func main() {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterModel(new(models.Category))
	orm.RegisterDataBase("default", "mysql","root:root@tcp(127.0.0.1:3333)/tx?charset=utf8",30)
	orm.RunSyncdb("default",false ,true )
	orm.Debug = true
	beego.Run()
}

