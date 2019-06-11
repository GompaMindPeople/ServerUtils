/*
@Time : 2019/5/24 9:47 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package controllers

import (
	"ServerUtils/models"
	"ServerUtils/service"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
)

type MainUtilController struct {
	beego.Controller
}

func (c *MainUtilController) URLMapping() {
	c.Mapping("getGameServerConfigForJson",c.GetGameServerConfigForJson)
	c.Mapping("listForButtonConf",c.ListForButtonConf)
	c.Mapping("saveAddDialog",c.SaveForButtonConf)
	c.Mapping("executeShell",c.ExecuteShell)
}

//@router /ManUtils/getGameServerConfigForJson [get]
func (c *MainUtilController) GetGameServerConfigForJson(){
	fmt.Print("getGameServerConfigForJson")
	c.ServeJSON()

}

//@router / [get]
func (c *MainUtilController) Index(){
	fmt.Print("getGameServerConfigForJson")
	c.TplName = "index.html"
}
//@router /listForButtonConf [get]
func (c *MainUtilController) ListForButtonConf(){
	c.Data["json"] = service.ListForButtonConf()
	c.ServeJSON()
}

//@router /saveAddDialog [post]
func (c *MainUtilController) SaveForButtonConf(){
	buttonName := c.GetString("buttonName")
	shell := c.GetString("shell")
	conf := models.ButtonConf{}
	//buttonName shell
	conf.Rows = make([]map[string]interface{},1)
	conf.Total = 1
	conf.Rows[0] = map[string]interface{}{}
	conf.Rows[0]["id"] = 1
	conf.Rows[0]["shell"] = shell
	conf.Rows[0]["buttonName"] = buttonName
	err := service.SaveForButtonConf(conf)
	if err != nil{
		c.Data["json"] = "发生错误"
	}
	c.ServeJSON()
}

//@router /deleteById [get]
func (c *MainUtilController) DeleteById(){
	id := c.GetString("id")
	i, _ := strconv.ParseFloat(id, 10)
	service.DeleteByIdForButtonConfig(i)
	c.ServeJSON()
}

//@router /executeShell [get]
func (c *MainUtilController) ExecuteShell(){
	shell := c.GetString("shell")
	buffer, _, e := service.SendShell(shell)
	if e != nil{
		logs.Error(e)
	}
	logs.Debug("executeShell--->",buffer.Bytes())
	c.Data["json"] = models.JsonMessage{Code: 200, Data: map[string]interface{}{"data": buffer.String()}}
	c.ServeJSON()
}



