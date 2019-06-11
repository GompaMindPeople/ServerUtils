/*
@Time : 2019/5/27 17:49 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package test

import (
	"ServerUtils/models"
	"ServerUtils/service"
	"log"
	"testing"
)

func TestButtonConfig(t *testing.T) {

	conf := models.ButtonConf{}

	//conf.Data = map[string]interface{}{}
	//conf.Index = 10
	//conf.Data["1"] = map[string] string{"1":"aaaa0"}
	//conf.Data["1"].(map[string]string)["122"] = "255"
	//conf.Data["2"] = map[string] string{"2":"aaaa"}
	//conf.Data["3"] = map[string] string{"3":"aaa2"}

	err := service.SaveForButtonConf(conf)
	log.Fatal(err)
}



