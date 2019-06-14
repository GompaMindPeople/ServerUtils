/*
@Time : 2019/6/10 15:08
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package server

import (
	"ServerUtils/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

/**
初始化mysql 驱动
*/
func MysqlInit() orm.Ormer {
	//注册驱动
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(err)
		return nil
	}
	//注册默认数据库
	err = orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/tool?charset=utf8")
	if err != nil {
		panic(err)
		return nil
	}
	//注册model
	orm.RegisterModel(new(models.SshConfig), new(models.Servers), new(models.ConfigAllTemplate), new(models.ServersTemplate))
	//自动建表
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		panic(err)
		return nil
	}
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	//初始化ormer
	newOrm := orm.NewOrm()
	return newOrm
}
