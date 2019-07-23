/*
@Time : 2019/6/24 16:20
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package models

//配置表,用于存储各种配置的数据.
type ConfigTable struct {
	Id          int        `orm:"column(Id);auto" description:"Id"`
	Type        *TypeTable `orm:"rel(fk)"`
	Group       int        `orm:"column(GroupName)" description:"Group,用于判断该行数据是否属于一组数据"`
	ConfigName  string     `orm:"column(ConfigName)" description:"ConfigName,配置的名称"`
	ConfigValue string     `orm:"column(ConfigValue)" description:"ConfigValue,配置的值"`
}

func (u *ConfigTable) TableName() string {
	return "t_config"
}
