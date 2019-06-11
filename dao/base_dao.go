/*
@Time : 2019/6/10 16:10 
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

type BaseDaoImpl struct {

}

type BaseDao interface {
	Insert(obj interface{})
	SelectList(obj *interface{})[]interface{}
}

/**
	根据model 插入 对应的 数据
 */
func (base  *BaseDaoImpl) Insert(obj interface{},num int){
	_, err := Global.OrmInstance.InsertMulti(num,obj)
	if err != nil{
		logs.Error(err)
		return
	}
}

/**
	从表中中查询一组数据,该数据按照 Map中的参数进行过滤数据
	key = 字段名字和 条件方式(详情参考beeGo orm高级查询)
	value = 所需要的匹配的值
 */
func (base *BaseDaoImpl) ListAll(tableName string,Map *map[string]interface{})[]orm.Params{
	var maps []orm.Params
	table := Global.OrmInstance.QueryTable(tableName)
	for  k,v :=  range *Map{
		table = table.Filter(k, v)
	}
	_, e := table.Values(&maps)
	if e != nil{
		logs.Error(e)
		return nil
	}
	return  maps
}


/**
	查询一组数据
 */
//func (base  *BaseDaoImpl) SelectList() []orm.ParamsList{
//	var list []orm.ParamsList
//	//emmm 直接返回所有列数据吧..影响不大.
//	_, err := Global.OrmInstance.QueryTable("")
//	if err != nil{
//		logs.Error(err)
//		return  nil
//	}
//	return list
//}


//
//var list []orm.ParamsList
////emmm 直接返回所有列数据吧..影响不大.
//_, err := Global.OrmInstance.Raw("SELECT * FROM `servers` WHERE `groupId` = ? ", groupId).ValuesList(&list)
//if err != nil{
//logs.Error(err)
//return  nil
//}