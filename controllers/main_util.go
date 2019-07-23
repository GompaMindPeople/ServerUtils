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
	c.Mapping("getGameServerConfigForJson", c.GetGameServerConfigForJson)
	c.Mapping("listForButtonConf", c.ListForButtonConf)
	c.Mapping("executeShell", c.ExecuteShell)
}

//@router /ManUtils/getGameServerConfigForJson [get]
func (c *MainUtilController) GetGameServerConfigForJson() {
	fmt.Print("getGameServerConfigForJson")
	c.ServeJSON()

}

//@router / [get]
func (c *MainUtilController) Index() {
	fmt.Print("getGameServerConfigForJson")
	c.TplName = "index.html"
}

//@router /listForButtonConf [get]
func (c *MainUtilController) ListForButtonConf() {
	service.ListForButtonConf()
	//c.Data["json"] =
	c.ServeJSON()
}

//@router /deleteById [get]
func (c *MainUtilController) DeleteById() {
	id := c.GetString("id")
	i, _ := strconv.ParseFloat(id, 10)
	service.DeleteByIdForButtonConfig(i)
	c.ServeJSON()
}

//@router /executeShell [get]
func (c *MainUtilController) ExecuteShell() {
	shell := c.GetString("shell")
	id := c.GetString("SSHId")
	i, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(err)
	}
	buffer, _, e := service.SendShell(shell, i)
	if e != nil {
		c.Data["json"] = models.JsonMessage{Code: 200, Data: map[string]interface{}{"data": e.Error()}}
		c.ServeJSON()
		return
	}
	c.Data["json"] = models.JsonMessage{Code: 200, Data: map[string]interface{}{"data": buffer.String()}}
	c.ServeJSON()
}

//@router /listGroupALL [get]
func (c *MainUtilController) ListGroupALL() {
	serviceImpl := service.ButtonGroupServiceImpl{}
	json := serviceImpl.ListAll()
	c.Data["json"] = json
	c.ServeJSON()
}

//@router /deleteGroupById [get]
func (c *MainUtilController) DeleteGroupById() {
	groupId := c.GetString("groupId")
	if groupId == "" {
		c.Data["json"] = "error"
		c.ServeJSON()
		return
	}
	i, e := strconv.Atoi(groupId)
	if e != nil {
		logs.Error(e)
		c.Data["json"] = "error"
		c.ServeJSON()
		return
	}
	serviceImpl := service.ButtonGroupServiceImpl{}
	serviceImpl.DeleteById(int64(i))
	c.ServeJSON()
}
