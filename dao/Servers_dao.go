/*
@Time : 2019/6/10 17:44 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package dao

import (
	"ServerUtils/Global"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type ServersDaoImpl struct {
	BaseDaoImpl
}

type ServersDao interface {
	//查询一组数据根据groupId
	ListServerByGroupId(groupId int)[]interface{}
}


/**
	查询一组服务器数据 根据 组Id
 */
func (server *ServersDaoImpl) ListServerByGroupId(groupId int)[]orm.Params{
	var maps []orm.Params
	table := Global.OrmInstance.QueryTable("servers")
	_, e := table.Filter("groupId", groupId).Values(&maps)
	if e != nil{
		logs.Error(e)
		return nil
	}
	return  maps
}

