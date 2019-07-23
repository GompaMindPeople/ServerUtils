/*
@Time : 2019/6/21 15:21
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package models

/**
存储的shell指令
*/
type ButtonShell struct {
	//buttonName":"asd","id":2,"shell":"fdsaf"
	Id         int          `orm:"column(id);auto" description:"ButtonShell"`
	ButtonName string       `orm:"column(buttonName);" description:"buttonName"`
	Shell      string       `orm:"column(shell);" description:"shell"`
	GroupId    *ButtonGroup `orm:"rel(fk)"`
}

func (u *ButtonShell) TableName() string {
	return "t_ButtonShell"
}
