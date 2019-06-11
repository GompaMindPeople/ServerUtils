/*
@Time : 2019/5/27 14:27 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package utils

import (
	"ServerUtils/models"
	"github.com/astaxie/beego/logs"
)

/**
	读取json格式的配置文件
 */
func ReadConfigFile(fileName string)(json map[string]interface{},err error){
	json, err = ReadJsonFileForMap("conf/" + fileName)
	return
}

func GetSSHConfig()models.SshConfig{
	config := models.SshConfig{}
	err := ReadJsonFileForStruct("conf/sshConfig.json",&config)
	if err != nil{
		logs.Error("读取sshConfig文件的报错",err)

	}
	return config
}


