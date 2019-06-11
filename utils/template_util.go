/*
@Time : 2019/6/11 9:53 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package utils

import (
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)


//读取某个模板
//模板路径在  resource/Template 下面
func ReadTemplate(fileName string)(content string){
	bytes, e := ioutil.ReadFile("resource/template/" + fileName)
	if e != nil{
		logs.Error(e)
		return content
	}
	content = string(bytes)
	return
}

//读取configAll.js模板文件的数据
func ReadConfigAllTemplate()(content string){
	return ReadTemplate("configAll_template.js")
}




