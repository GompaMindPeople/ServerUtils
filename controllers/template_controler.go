/*
@Time : 2019/6/12 15:31
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package controllers

import (
	"ServerUtils/models"
	"ServerUtils/service"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type TemplateController struct {
	beego.Controller
}

//模板服务的实体
var impl = service.TemplateServiceImpl{}

//@router /listServerName [get]
func (t *TemplateController) ListServerName() {
	data := impl.ListTemplateForServerName("t_configalltemplate")
	arr := make([]interface{}, 0)
	for _, v := range data.([]orm.ParamsList) {
		temp := map[string]interface{}{}
		temp["TemplateId"] = v[0]
		temp["ServerName"] = v[1]
		if temp != nil {
			arr = append(arr, temp)
		}
	}
	t.Data["json"] = arr
	t.ServeJSON()

}

//@router /getTemplateById [get]
func (t *TemplateController) GetTemplateById() {
	s := t.GetString("TemplateId")
	i, e := strconv.Atoi(s)
	if e != nil {
		logs.Error(e)
	}
	data := impl.GetTemplateById(int64(i))
	t.Data["json"] = data
	t.ServeJSON()
}

//@router /saveConfigAll [post]
func (t *TemplateController) SaveConfigAll() {
	template := models.ConfigAllTemplate{}
	if err := t.ParseForm(&template); err != nil {
		logs.Error("ParseFormToTemplate error---->", err)
	}
	impl.SaveConfigAllTemplate(&template)
	t.ServeJSON()
}

//@router /saveServers [post]
func (t *TemplateController) SaveServers() {
	template := models.ServersTemplate{}
	if err := t.ParseForm(&template); err != nil {
		logs.Error("ParseFormToTemplate error---->", err)
	}
	impl.SaveServersTemplate(template)
	t.ServeJSON()
}

//@router /generateConfigAllTemplate [get]
func (t *TemplateController) GenerateConfigAllTemplate() {
	s := t.GetString("templateId")
	var i int
	var e error
	if s != "" {
		i, e = strconv.Atoi(s)
		if e != nil {
			logs.Error(e)
			return
		}
	}
	impl.WriteConfigAll(i)
	t.ServeJSON()
}

//下载ConfigAll
//@router /downloadConfigAllTemplate [get]
func (t *TemplateController) DownloadConfigAllTemplate(fileName string) {
	//第一个参数是文件的地址，第二个参数是下载显示的文件的名称
	t.Ctx.Output.Download("resource/configAll.js")
}
