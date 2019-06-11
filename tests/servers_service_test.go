/*
@Time : 2019/6/10 19:18 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package test

import (
	"ServerUtils/service"
	"testing"
)

func TestWrite(t *testing.T) {

	impl := service.TemplateServiceImpl{}
	impl.WriteServers(1)
}