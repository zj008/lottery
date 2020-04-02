package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "lottery/huancai/routers"
)

func init() {
	mysqlhost := beego.AppConfig.String("mysqlhost")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqldb   := beego.AppConfig.String("mysqldb")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqluser, mysqlpass, mysqlhost, mysqlport, mysqldb)
	//mysql := mysqluser + ":" + mysqlpass + "@/" + mysqlhost + ":" + mysqlport + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", dataSource)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetStaticPath("/", "views")
	beego.SetStaticPath("/", "static")
	orm.Debug = true
	beego.Run()
}
