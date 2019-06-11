/*
@Time : 2019/6/10 18:24 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package service

import (
	"ServerUtils/dao"
	"ServerUtils/utils"
	"ServerUtils/utils/stringutils"
	"github.com/astaxie/beego/logs"
)

type TemplateService interface {
	WriteServers(groupId int)
	WriteConfigAll(templateId int)
}

type TemplateServiceImpl struct {

}



//按照模板写出 一份 configAll.js文件,并需要将该文件进行持久化
func (s *TemplateServiceImpl) WriteConfigAll(templateId int){
	Map := map[string]interface{}{}
	Map["templateId"] = templateId
	impl := dao.BaseDaoImpl{}
	all := impl.ListAll("t_ConfigAllTemplate", &Map)
	content := utils.ReadConfigAllTemplate()
	stringutils.ReplacePlaceholders(&content,all[0])
	writeConfigAllFile(content)

}


//写..哈...tui
func (s *TemplateServiceImpl) WriteServers(groupId int){
	impl := dao.ServersDaoImpl{}
	list := impl.ListServerByGroupId(groupId)
	writeServersFile(list)
	logs.Debug(list)

}


/**
	写出servers格式的文件
 */
func writeServersFile(Map interface{}){
	utils.WriteJson("resource/servers.json",Map)
}

/**
	写出servers格式的文件
 */
func writeConfigAllFile(Map interface{}){
	utils.WriteJson("resource/configAll.js",Map)
}