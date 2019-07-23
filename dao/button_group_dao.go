/*
@Time : 2019/6/21 15:34
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package dao

import (
	"ServerUtils/models"
	"github.com/astaxie/beego/logs"
)

type ButtonGroupImpl struct {
	BaseDaoImpl
}

type ButtonGrouper interface {
	//查询一个组的数据,根据Id
	GetGroupById(id int) models.ButtonGroup
	//查询所有的List
	ListGroup() []models.ButtonGroup
	//插入一组数据
	InertGroup(group *models.ButtonGroup)
	//更新 数据
	UpdateGroup(group *models.ButtonGroup)
}

//根据Id查询一个数据
func (bg *ButtonGroupImpl) GetGroupById(id int64) models.ButtonGroup {
	group := models.ButtonGroup{GroupId: id}
	//TableName := group.TableName()
	bg.GetByConditions(&group, "GroupId")
	return group
}

//根据Id查询一个数据
func (bg *ButtonGroupImpl) ListGroup() []models.ButtonGroup {
	var bgs []models.ButtonGroup
	group := models.ButtonGroup{}
	_, err := bg.ListForAll(&bgs, group.TableName(), nil)
	//TableName := group.TableName()
	if err != nil {
		logs.Error(err)
	}
	return bgs
}

func (bg *ButtonGroupImpl) InertGroup(group *models.ButtonGroup) {
	bg.InsertOne(group)
}

func (bg *ButtonGroupImpl) UpdateGroup(group *models.ButtonGroup) {
	bg.UpdateOne(group)
}
