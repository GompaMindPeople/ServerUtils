/*
@Time : 2019/6/24 17:50
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package service

import (
	"ServerUtils/dao"
	"ServerUtils/models"
)

type TypeServiceer interface {
	WriteServers(groupId int)
	WriteConfigAll(templateId int)
	GetTemplateById(TemplateId int64) interface{}
}

type TypeServiceImpl struct {
}

var impl = dao.TypeImpl{}

func (tsi *TypeServiceImpl) GetTypeById(id int) models.TypeTable {
	obj := impl.GetTypeById(id)
	return obj
}
