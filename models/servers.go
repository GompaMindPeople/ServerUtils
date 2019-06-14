/*
@Time : 2019/6/10 17:16
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package models

//服务器组列表 用于 servers.json 的配置
type Servers struct {
	Sid int `orm:"column(Sid);auto" description:"Sid"`
	//服务器组Id,用来区分 数据属于某一组 服务器
	GroupId        int    `orm:"column(groupId)" description:"服务器组Id,用来区分数据"`
	RowName        string `orm:"column(rowName)" `
	Order          int    `orm:"column(order)" `
	Id             string `orm:"column(id)" `
	Host           string `orm:"column(host)" `
	ClientPort     int    `orm:"column(clientPort)" `
	Port           int    `orm:"column(port)" `
	Frontend       bool   `orm:"column(frontend)" `
	Db             string `orm:"column(db)" `
	MaxConnections string `orm:"column(maxConnections)"`
}

func (u *Servers) TableName() string {
	return "t_servers"
}
