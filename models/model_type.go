/*
@Time : 2019/6/24 16:13
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package models

//类型表.抽象的一张表,用于存储各种类型数据
type TypeTable struct {
	Id       int    `orm:"column(Id);auto" description:"Id"`
	TypeName string `orm:"column(TypeName);" description:"类型的名称"`
	Note     string `orm:"column(Note);" description:"类型的备注,用于说明每个类型的作用."`
}

func (u *TypeTable) TableName() string {
	return "t_type"
}
