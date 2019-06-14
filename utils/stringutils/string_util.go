/*
@Time : 2019/6/11 10:05
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package stringutils

import (
	"github.com/astaxie/beego/logs"
	"reflect"
	"strings"
)

/**
填充占位符数据
*/
func ReplacePlaceholders(content *string, Map map[string]interface{}) {
	for k, v := range Map {
		if reflect.TypeOf(v).String() == "string" {
			*content = strings.Replace(*content, "#"+k+"#", v.(string), -1)
		}
	}
	logs.Debug(content)
	return
}
