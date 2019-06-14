/*
@Time : 2019/6/10 18:24
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package service

import (
	"ServerUtils/dao"
	"ServerUtils/models"
	"ServerUtils/utils"
	"ServerUtils/utils/stringutils"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

type TemplateService interface {
	WriteServers(groupId int)
	WriteConfigAll(templateId int)
	GetTemplateById(TemplateId int64) interface{}
}

type TemplateServiceImpl struct {
}

func (s *TemplateServiceImpl) ListTemplateForServerName(tableName string) interface{} {
	impl := dao.BaseDaoImpl{}
	values := impl.ListForFieldValue(tableName, "TemplateId", "ServerName")
	return values
}

//保存一个 servers模板
func (s *TemplateServiceImpl) SaveServersTemplate(obj models.ServersTemplate) {
	SaveTemplate(obj)
}

//保存一个 ConfigAll.js模板
func (s *TemplateServiceImpl) SaveConfigAllTemplate(obj *models.ConfigAllTemplate) {
	if obj.TemplateId != 0 {
		UpdateTemplate(obj)
		return
	}
	SaveTemplate(obj)
}

func UpdateTemplate(bean interface{}) {
	impl := dao.BaseDaoImpl{}
	impl.UpdateOne(bean)
}

//保存模板
func SaveTemplate(bean interface{}) {
	impl := dao.BaseDaoImpl{}
	impl.InsertOne(bean)
}

/**
根据模板的Id,获取模板数据
*/
func (s *TemplateServiceImpl) GetTemplateById(TemplateId int64) interface{} {
	Map := map[string]interface{}{}
	Map["templateId"] = TemplateId
	impl := dao.BaseDaoImpl{}
	all := impl.ListAll("t_configalltemplate", &Map)
	if len(all) > 0 {
		return all[0]
	}
	return all
}

//按照模板写出 一份 configAll.js文件,并需要将该文件进行持久化
func (s *TemplateServiceImpl) WriteConfigAll(templateId int) {
	Map := map[string]interface{}{}
	Map["templateId"] = templateId
	impl := dao.BaseDaoImpl{}
	all := impl.ListAll("t_configalltemplate", &Map)
	content := utils.ReadConfigAllTemplate()
	stringutils.ReplacePlaceholders(&content, all[0])
	writeConfigAllFile(&content)

}

//写............
func (s *TemplateServiceImpl) WriteServers(templateId int) {
	Map := map[string]interface{}{}
	Map["ServersTemplateId"] = templateId
	impl := dao.BaseDaoImpl{}
	all := impl.ListAll("t_servers_template", &Map)
	content := utils.ReadServersTemplate()
	stringutils.ReplacePlaceholders(&content, all[0])
	writeServersFile(&content)
}

//写..哈...tui
//func (s *TemplateServiceImpl) WriteServers(groupId int){
//	impl := dao.ServersDaoImpl{}
//	list := impl.ListServerByGroupId(groupId)
//	writeServersFile(list)
//	logs.Debug(list)
//
//}

/**
写出servers格式的文件
*/
func writeServersFile(content *string) {
	writeResourceFile("servers.json", content)
}

/**
写出servers格式的文件
*/
func writeConfigAllFile(content *string) {
	writeResourceFile("configAll.js", content)
}

func writeResourceFile(fileName string, content *string) {
	err := ioutil.WriteFile("resource/"+fileName, []byte(*content), 0666)
	if err != nil {
		logs.Error(err)
	}
}
