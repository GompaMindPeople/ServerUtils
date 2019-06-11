/*
@Time : 2019/5/27 11:10 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package service

import (
	"ServerUtils/models"
	"ServerUtils/server"
	"bytes"
	"errors"
	"github.com/astaxie/beego/logs"
)


func SendShell(cmd string)(*bytes.Buffer,*bytes.Buffer,error){
	var	stdout ,stderr  bytes.Buffer
	if cmd == ""{
		return &stdout,&stderr,errors.New("指令为空")
	}
	//session, err := server.GetSSHSession()
	bean := server.SshBean{}

	bean.RunSsh(cmd,&stdout,&stderr)

	//if err != nil{
	//	return  &stdOut,&stdErr,nil
	//}
	//session.SendMessage(cmd,&stdOut,&stdErr)
	return  &stdout,&stderr,nil
}




/**
	获取ssh的配置数据
 */
func GetSSHConfig() *models.SshConfig{
	config := models.SshConfig{}
	one := config.SelectOne()
	return one
}
/**
	保存ssh配置信息.
	有坑!
 */
func SaveSSHConfig(config *models.SshConfig){
	i, e := config.Update()
	if e!=nil{
		logs.Error(e)
		return
	}
	if i == 0{
		_, e := config.InsertOne()
		if e != nil{
			logs.Error(e)
		}
	}
}





