/*
@Time : 2019/6/21 16:30
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package service

import (
	"ServerUtils/dao"
	"ServerUtils/models"
	"github.com/astaxie/beego/logs"
)

type ButtonGroupServiceImpl struct {
}

var imp = dao.ButtonGroupImpl{}

//查询所有的 组数据
func (bsi *ButtonGroupServiceImpl) ListAll() (model models.EasyUIModel) {
	groups := imp.ListGroup()
	model.Rows = groups
	return
}

//插入一条 数据
func (bsi *ButtonGroupServiceImpl) Insert(group models.ButtonGroup) {
	imp.InertGroup(&group)
	return
}

//删除 一个数据工具Id
func (bsi *ButtonGroupServiceImpl) DeleteById(groupId int64) {
	buttonGroup := models.ButtonGroup{}
	buttonGroup.GroupId = groupId
	_, err := imp.Delete(&buttonGroup, "GroupId")
	if err != nil {
		logs.Error(err)
	}
	return
}
