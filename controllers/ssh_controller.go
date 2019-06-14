/*
@Time : 2019/5/27 11:47
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package controllers

import (
	"ServerUtils/models"
	"ServerUtils/service"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type SSHController struct {
	beego.Controller
}

//
func (ssh *SSHController) URLMapping() {
	ssh.Mapping("getSshConfigById", ssh.GetSshConfig)
	ssh.Mapping("saveSSHConfig", ssh.SaveSSHConfig)
}

//@router /getSshConfigById [get]
func (ssh *SSHController) GetSshConfig() {
	s := ssh.GetString("SSHId")
	var i int
	var e error
	if s != "" {
		i, e = strconv.Atoi(s)
		if e != nil {
			logs.Error(e)
		}
	}
	config := service.GetSSHConfig(i)
	fmt.Print("getSshConfig ")
	ssh.Data["json"] = config
	ssh.ServeJSON()
}

//@router /saveSSHConfig [post]
func (ssh *SSHController) SaveSSHConfig() {
	SSHId := ssh.GetString("SSHId")
	userName := ssh.GetString("UserName")
	password := ssh.GetString("Password")
	hostName := ssh.GetString("HostName")
	port := ssh.GetString("Port")
	i, e := strconv.Atoi(port)
	var id int
	if SSHId != "" {
		id, e = strconv.Atoi(SSHId)
	}

	if e != nil {
		ssh.Data["json"] = errors.New("错误的端口号")
		logs.Error(e, "-->错误的端口号")
		ssh.ServeJSON()
		return
	}

	config := models.SshConfig{Id: id, UserName: userName, Password: password, HostName: hostName, Port: int16(i)}
	service.SaveSSHConfig(&config)
	ssh.ServeJSON()
}

//@router /listSSHCombobox [get]
func (ssh *SSHController) ListSSHCombobox() {
	combobox := service.ListSSHCombobox()
	arr := make([]interface{}, 0)
	for _, v := range combobox.([]orm.ParamsList) {
		temp := map[string]interface{}{}
		temp["SSHId"] = v[0]
		temp["HostName"] = v[1]
		if temp != nil {
			arr = append(arr, temp)
		}
	}
	ssh.Data["json"] = arr
	ssh.ServeJSON()

}

//func (c *SSHController) URLMapping() {
//	c.Mapping("sendShell",c.SendShell)
//}
////@router /ssh/sendShell:shellString [get]
//func (c *SSHController) SendShell(){
//	var err error
//	cmd := c.GetString("shellString")
//	stdOut,_, err := service.SendShell(cmd)
//	if err != nil{
//		log.Fatal("执行sendShell发生错误-->",err)
//		c.Data["json"] = models.JsonMessage{Code: 500, Data: map[string]interface{}{"error": err}}
//	}else{
//		c.Data["json"] = models.JsonMessage{Code: 200, Data: map[string]interface{}{"data": stdOut.Bytes()}}
//	}
//	c.ServeJSON()
//}
