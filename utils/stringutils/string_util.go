/*
@Time : 2019/6/11 10:05 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package stringutils

import (
	"github.com/astaxie/beego/logs"
	"strings"
)
/**
	填充占位符数据
 */
func ReplacePlaceholders(content *string,Map map[string]interface{}){
	for k,v := range Map {
		strings.Replace(*content,"#"+k+"#",v.(string),-1)
	}
	logs.Debug(content)
	return
}

