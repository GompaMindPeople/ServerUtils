/*
@Time : 2019/5/24 16:11 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package utils

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io"
	"log"
	"os"
)
/**
	写出json文件
 */
func WriteJson(fileName string,bean interface{}){
	if fileName == "" || bean == nil{
		logs.Error("文件名和实体不能为空!")
		return
	}
	//开始返回一个文件的配置和信息,这边主要是用来判断文件是否存在
	_, e := os.Stat(fileName)
	var (
		filePtr *os.File
		err error
	)
	if os.IsNotExist(e){
		filePtr, err = os.Create(fileName)
	}else{
		filePtr, err = os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0666)
	}
	if err != nil{
		log.Fatal("OS.Open方式调用失败-->",err)
		return
	}
	//生成文件编码器
	encoder:= json.NewEncoder(filePtr)
	//使用编码器将结构体编码到文件中
	err = encoder.Encode(bean)
	if err!=nil{
		fmt.Println("文件写入文件失败！",err)
	}
	defer filePtr.Close()
}
/**
	读json 文件
 */
func ReadJsonFileForMap(fileName string)(map[string]interface{}, error){
	var jsonMap = map[string] interface{}{}
	filePtr, err := os.Open(fileName)
	if err != nil{
		return  nil,err
	}
	decodeForJsonString(filePtr,&jsonMap)
	defer filePtr.Close()
	return jsonMap,nil
}

//对于json解码的封装
func decodeForJsonString(filePtr io.Reader,obj interface{}){
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err := decoder.Decode(obj)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())
	} else {
		fmt.Println("Decoder success")
	}

}
//读json文件并转换成对应的struct
func ReadJsonFileForStruct(fileName string,obj interface{})error{
	filePtr, err := os.Open(fileName)
	if err != nil{
		return  err
	}
	decodeForJsonString(filePtr,&obj)
	defer filePtr.Close()
	return nil
}




