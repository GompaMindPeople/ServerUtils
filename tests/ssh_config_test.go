/*
@Time : 2019/6/10 15:35
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package test

import (
	"ServerUtils/Global"
	"ServerUtils/models"
	"ServerUtils/server"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"testing"
)

func init() {
	fmt.Println("初始化测试用例")
	Global.OrmInstance = server.MysqlInit()

}

func TestInsertOne(t *testing.T) {
	//"userName": "root",
	//	"password": "2098231.xzm",
	//	"hostName": "192.168.5.189",
	//	"port": 22

	config := models.SshConfig{UserName: "root", Password: "2098231.xzm", HostName: "192.168.5.189", Port: 22}
	i, e := config.InsertOne()
	if e != nil {
		logs.Error(e)
		return
	}
	t.Log(i)
}

func TestSelectOne(t *testing.T) {
	//config := models.SshConfig{}
	//one := config.SelectOne()
	//t.Log(one)
}

func TestUpdate(t *testing.T) {
	config := models.SshConfig{UserName: "root", Password: "2098231.xzm", HostName: "192.168.5.189", Port: 22}
	i, _ := config.Update()
	t.Log(i)
}
