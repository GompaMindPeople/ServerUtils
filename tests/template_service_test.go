/*
@Time : 2019/6/11 15:43
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package test

import (
	"ServerUtils/service"
	"testing"
)

func TestTemplateServiceImpl_WriteConfigAll(t *testing.T) {
	//impl := dao.BaseDaoImpl{}
	//impl.ListAll()
	impl := service.TemplateServiceImpl{}
	impl.WriteConfigAll(1)
}
