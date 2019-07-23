package main

import (
	"ServerUtils/Global"
	_ "ServerUtils/routers"
	"ServerUtils/server"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func main() {

	//初始化 mysql 的驱动,连接mysql......嘤嘤嘤 如果有需求在修改
	Global.OrmInstance = server.MysqlInit()
	//if beego.AppConfig.String("PprofOn")  {
	//	beego.Reg(`/debug/pprof`, &ProfController{})
	//	beego.RegisterController(`/debug/pprof/:pp([\w]+)`, &ProfController{})
	//}
	//
	//beego.Pr
	beego.Run()
}
