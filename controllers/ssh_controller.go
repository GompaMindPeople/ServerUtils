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
	"strconv"
)

type SSHController struct {
	beego.Controller
}

//
func (ssh *SSHController) URLMapping() {
	ssh.Mapping("getSshConfig",ssh.GetSshConfig)
	ssh.Mapping("saveSSHConfig",ssh.SaveSSHConfig)
}


//@router /getSshConfig [get]
func (ssh *SSHController) GetSshConfig(){
	config := service.GetSSHConfig()
	fmt.Print("getSshConfig ")
	ssh.Data["json"] = config
	ssh.ServeJSON()
}

//@router /saveSSHConfig [post]
func (ssh *SSHController) SaveSSHConfig(){
	userName := ssh.GetString("userName")
	password := ssh.GetString("password")
	hostName := ssh.GetString("hostName")
	port := ssh.GetString("port")
	 i, e := strconv.Atoi(port)
	if e != nil{
		ssh.Data["json"] = errors.New("错误的端口号")
		logs.Error(e,"-->错误的端口号")
		ssh.ServeJSON()
		return
	}
	config := models.SshConfig{UserName: userName, Password: password, HostName: hostName, Port: int16(i)}
	service.SaveSSHConfig(&config)
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

