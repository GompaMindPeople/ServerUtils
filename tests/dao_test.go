/*
@Time : 2019/6/10 18:46 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package test

import (
	"ServerUtils/dao"
	"ServerUtils/models"
	"testing"
)

func TestInsert(t *testing.T){
	impl := dao.ServersDaoImpl{}
	servers := models.Servers{GroupId: 1, RowName: "gate", Order: 1, Id: "gate", Host: "host1997", ClientPort: 31, Port: 30, Frontend: true, Db: "db"}
	arr := []models.Servers{servers}
	impl.Insert(&arr,1)
	t.Log("2333")
}
func TestListServerByGroupId(t *testing.T){
	impl := dao.ServersDaoImpl{}
	list := impl.ListServerByGroupId(1)
	t.Log("2333",list)
}


func TestInsertTemplate(t *testing.T){
	template := models.ConfigAllTemplate{ServerId: "asd", ServerHttpPort: "dsa",
		ServerIp: "123", ServerName: "123", ServerOpenDate: "321", ServerVersion: "321", ServerPlatform: "421", ServerPartner: "333",
		RechargeServerPort: "asd", RechargeServerIp: "asd",
		AcctServerPort: "asd", AcctServerIp: "asd",
		MongoUser: "fdsf", MongoPwd: "sdf", MongoHost: "sF",
		ChatDBUser: "SDF", ChatDBPwd: "SDFSD", ChatDBHost: "GS",
		MysqlHost: "SD", MysqlPort: "DFS", MysqlDataBase: "ASDF", MysqlUser: "SDFAS", MysqlPassword: "SDF"}
	arr := []models.ConfigAllTemplate{template}
	impl := dao.ServersDaoImpl{}
	 impl.Insert(&arr,1)
}
