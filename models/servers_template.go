/*
@Time : 2019/6/11 11:33
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package models

//servers的配置
type ServersTemplate struct {
	ServersTemplateId   int    `orm:"column(ServersTemplateId);auto" description:"ServersTemplateId"`
	GateClientPort      string `orm:"column(GateClientPort);" description:"GateClientPort"`
	GatePort            string `orm:"column(GatePort);" description:"GatePort"`
	ConnectorPort       string `orm:"column(ConnectorPort);" description:"ConnectorPort"`
	ConnectorClientHost string `orm:"column(ConnectorClientHost);" description:"ConnectorClientHost"`
	ConnectorClientPort string `orm:"column(ConnectorClientPort);" description:"ConnectorClientPort"`
	WorldPort           string `orm:"column(WorldPort);" description:"WorldPort"`
	AreaPort            string `orm:"column(AreaPort);" description:"AreaPort"`
	WarClientHost       string `orm:"column(WarClientHost);" description:"WarClientHost"`
	WarClientPort       string `orm:"column(WarClientPort);" description:"WarClientPort"`
}

func (u *ServersTemplate) TableName() string {
	return "t_servers_template"
}
