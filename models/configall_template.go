/*
@Time : 2019/6/11 11:33 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package models





type ConfigAllTemplate struct {
	//本地服务器的配置
	TemplateId int `orm:"column(templateId);auto" description:"templateId"`
	ServerId  string `orm:"column(serverId);" description:"serverId"`
	ServerHttpPort string `orm:"column(serverHttpPort);" description:"serverHttpPort"`
	ServerIp string `orm:"column(serverIp);" description:"serverIp"`
	ServerName string `orm:"column(serverName);" description:"serverName"`
	ServerOpenDate string `orm:"column(serverOpenDate);" description:"serverOpenDate"`
	ServerVersion string `orm:"column(serverVersion);" description:"serverVersion"`
	ServerPlatform string `orm:"column(serverPlatform);" description:"serverPlatform"`
	ServerPartner  string `orm:"column(serverPartner);" description:"serverPartner"`
	//帐号,充值的配置
	RechargeServerPort   string `orm:"column(rechargeServerPort);" description:"rechargeServerPort"`
	RechargeServerIp  string `orm:"column(rechargeServerIp);" description:"rechargeServerIp"`
	AcctServerPort   string `orm:"column(acctServerPort);" description:"acctServerPort"`
	AcctServerIp  string `orm:"column(acctServerIp);" description:"acctServerIp"`
	//mongo的配置
	MongoUser  string `orm:"column(mongoUser);" description:"mongoUser"`
	MongoPwd  string `orm:"column(mongoPwd);" description:"mongoPwd"`
	MongoHost  string `orm:"column(mongoHost);" description:"mongoHost"`
	ChatDBUser  string `orm:"column(chatDBUser);" description:"chatDBUser"`
	ChatDBPwd  string `orm:"column(chatDBPwd);" description:"chatDBPwd"`
	ChatDBHost  string `orm:"column(chatDBHost);" description:"chatDBHost"`
	//mysql 的配置
	MysqlHost string `orm:"column(mysqlHost);" description:"mysqlHost"`
	MysqlPort string `orm:"column(mysqlPort);" description:"mysqlPort"`
	MysqlDataBase string `orm:"column(mysqlDataBase);" description:"mysqlDataBase"`
	MysqlUser string `orm:"column(mysqlUser);" description:"mysqlUser"`
	MysqlPassword string `orm:"column(mysqlPassword);" description:"mysqlPassword"`

}

func (u *ConfigAllTemplate) TableName() string {
	return "t_ConfigAllTemplate"
}