/*
@Time : 2019/5/24 16:39 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package test

import (
	"ServerUtils/models"
	"ServerUtils/utils"
	"fmt"
	"log"
	"testing"
)

//定义配置文件解析后的结构
type MongoConfig struct {
	MongoAddr      string
	MongoPoolLimit int
	MongoDb        string
	MongoCol       string
}

type Config struct {
	Addr  string
	Mongo MongoConfig
}
type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}


type test struct {
	a string
}
func TestReadJsonTest(t *testing.T) {
	//strings := make(map[string]string)
	strings, e := utils.ReadJsonFileForMap("conf/test.json")
	if e != nil{
		log.Fatal(e)
	}
	fmt.Print("----<>",strings)
}


func TestWriteJson(t *testing.T){
	 utils.WriteJson("conf/test1.json",MongoConfig{"123",1,"aaa","123"})
}




func TestReadJsonFileForStruct(t *testing.T){
	config := models.SshConfig{}
	err := utils.ReadJsonFileForStruct("conf/sshConfig.json", &config)
	if err != nil{
		t.Log(err)
	}
	log.Print(config)
}
