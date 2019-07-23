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
	"reflect"
)

type BaseDaoImpl struct {
}

type BaseDao interface {
	Insert(obj interface{})
	SelectList(obj *interface{}) []interface{}
	InsertOne(obj interface{})
	UpdateOne(obj interface{})
}

/**
更新一组数据
*/
func (base *BaseDaoImpl) UpdateOne(obj interface{}) {
	_, err := Global.OrmInstance.Update(obj)
	if err != nil {
		logs.Error(err)
		return
	}
}

/**
根据model 插入 一组 数据
*/
func (base *BaseDaoImpl) Insert(obj interface{}, num int) {
	_, err := Global.OrmInstance.InsertMulti(num, obj)
	if err != nil {
		logs.Error(err)
		return
	}
}

//插入一个数据到数据库
func (base *BaseDaoImpl) InsertOne(obj interface{}) {
	_, err := Global.OrmInstance.Insert(obj)
	if err != nil {
		logs.Error(err)
		return
	}
}

//查询一个字段的所有列表数据
func (base *BaseDaoImpl) ListForFieldValue(tableName, TemplateId, FieldName string) interface{} {
	var list []orm.ParamsList
	////emmm 直接返回所有列数据吧..影响不大.
	_, err := Global.OrmInstance.Raw("SELECT " + TemplateId + "," + FieldName + " FROM " + tableName + " ").ValuesList(&list)
	if err != nil {
		logs.Error("ListForFieldValue--Select Error-->", err)
		return nil
	}
	return list
}

/**
删除数据,根据obj对应的数据进行操作
*/
func (base *BaseDaoImpl) Delete(obj interface{}, str string) (id int64, err error) {
	id, err = Global.OrmInstance.Delete(obj, str)
	return
}

/**
从表中中查询一组数据,该数据按照 Map中的参数进行过滤数据
key = 字段名字和 条件方式(详情参考beeGo orm高级查询)
value = 所需要的匹配的值
------------------弃用-----------------
*/
func (base *BaseDaoImpl) ListAll(tableName string, Map *map[string]interface{}) []orm.Params {
	var maps []orm.Params
	table := fillFilter(tableName, Map)
	_, e := table.Values(&maps)
	if e != nil {
		logs.Error(e)
		return nil
	}
	return maps
}

/**
根据一个条件进行查询,查询一个对象
*/
func (base *BaseDaoImpl) GetByConditions(obj interface{}, conditions string) {
	//TableName := group.TableName()
	err := Global.OrmInstance.Read(obj, conditions)
	if err != nil {
		logs.Error(err)
	}
}

/**
根据一个对象对应的属性名称 和 所对应的值进行 查询 .返回 该对象
*/
func (base *BaseDaoImpl) GetByConditionsOne(obj interface{}, conditions, value string) {
	//自动填充 对于
	SetObjAttrByString(obj, conditions, value)
	err := Global.OrmInstance.Read(obj, conditions)
	if err != nil {
		logs.Error(err)
	}
}

/**
从表中中查询一组数据,该数据按照 Map中的参数进行过滤数据
list 需要通过指针 传入一组缓存数据,用于填充数据
tableName  表名
key = 字段名字和 条件方式(详情参考beeGo orm高级查询)
value = 所需要的匹配的值
*/
func (base *BaseDaoImpl) ListForAll(list interface{}, tableName string, Map *map[string]interface{}) (num int64, err error) {
	table := fillFilter(tableName, Map)
	i, err := table.All(list)
	return i, err
}

//根据表名和 Map条件进行过滤数据.
//key为 字段,value为需要过滤的数据
func fillFilter(tableName string, Map *map[string]interface{}) orm.QuerySeter {
	table := Global.OrmInstance.QueryTable(tableName)
	if Map == nil {
		return table
	}
	for k, v := range *Map {
		table = table.Filter(k, v)
	}
	return table
}

func SetObjAttrByString(obj interface{}, key, val string) {
	valueOf := reflect.ValueOf(obj)
	elem := valueOf.Elem()
	name := elem.FieldByName(key)
	name.SetString(val)
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
