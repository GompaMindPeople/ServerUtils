/*
@Time : 2019/6/21 15:24
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package models

/**
按钮的组名-->主要用于 给按钮进行分组
*/
type ButtonGroup struct {
	GroupId   int64  `orm:"column(groupId);auto" description:"ButtonGroupId"`
	GroupName string `orm:"column(groupName);" description:"GroupName"`
}

func (u *ButtonGroup) TableName() string {
	return "t_ButtonGroup"
}
