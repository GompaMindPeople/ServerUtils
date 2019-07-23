/*
@Time : 2019/6/21 15:31
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package dao

type ButtonShellImpl struct {
	BaseDaoImpl
}

type ButtonSheller interface {
	//查询一组数据根据groupId
	ListServerByGroupId(groupId int) []interface{}
}
