/*
@Time : 2019/6/24 16:36
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package dao

import (
	"ServerUtils/models"
	"github.com/astaxie/beego/logs"
)

type TypeImpl struct {
	BaseDaoImpl
}

type Typeer interface {
	//查询一组数据根据TypeName
	GetTypeById(id int) models.TypeTable

	//查询一组数据根据TypeName
	GetTypeByName(typeName string) models.TypeTable

	//查询所有数据
	ListTypeAll() []models.TypeTable
	//更新类型根据Id, emmmm其实就是将model中的Id设置对应的值即可
	UpdateTypeById(table models.TypeTable) int64

	DeleteTypeByName(name string)

	DeleteTypeById(id int)
}

func (ti *TypeImpl) DeleteTypeById(id int) {
	table := models.TypeTable{}
	table.Id = id
	_, err := ti.Delete(&table, "Id")
	if err != nil {
		logs.Error(err)
	}
}

func (ti *TypeImpl) DeleteTypeByName(name string) {
	table := models.TypeTable{}
	table.TypeName = name
	_, err := ti.Delete(&table, "TypeName")
	if err != nil {
		logs.Error(err)
	}
}

func (ti *TypeImpl) GetTypeById(id int) models.TypeTable {
	table := models.TypeTable{}
	table.Id = id
	ti.GetByConditions(&table, "Id")
	return table
}

func (ti *TypeImpl) GetTypeByName(typeName string) models.TypeTable {
	table := models.TypeTable{}
	ti.GetByConditionsOne(&table, "TableName", typeName)
	return table
}

func (ti *TypeImpl) ListTypeAll() []models.TypeTable {
	var tables = make([]models.TypeTable, 10)
	table := models.TypeTable{}
	_, err := ti.ListForAll(&tables, table.TableName(), nil)
	if err != nil {

		logs.Error(err)
		return nil
	}
	return tables
}
