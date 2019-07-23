/*
@Time : 2019/5/27 17:04
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package service

import (
	"ServerUtils/utils"
	"fmt"
)

func ListForButtonConf() (map[string]interface{}, int) {
	Map, _ := utils.ReadJsonFileForMap("conf/ButtonConf.json")
	return Map, 0
}

//func  SaveForButtonConf(conf models.ButtonConf)error{
//	if &conf == nil{
//		return errors.New("ButtonConf对象为空")
//	}
//
//	//forStruct := utils.ReadJsonFileForStruct("conf/ButtonConf.json", conf)
//	Map, _ := utils.ReadJsonFileForMap("conf/ButtonConf.json")
//	if Map != nil{
//		index := Map["Total"].(float64)
//		i := int (index)
//		i++
//		var data  = Map["Rows"].([]interface{})
//		conf.Rows[0]["id"] = i
//		data = append(data, conf.Rows[0])
//		Map["Total"] =i
//		Map["Rows"] = data
//		utils.WriteJson("conf/ButtonConf.json",Map)
//		return nil
//	}
//	//fmt.Print(forStruct)
//	utils.WriteJson("conf/ButtonConf.json",conf)
//	return nil
//}

func DeleteByIdForButtonConfig(id float64) {
	Map, _ := utils.ReadJsonFileForMap("conf/ButtonConf.json")
	if Map != nil {
		var data = Map["Rows"].([]interface{})
		for k, v := range data {
			fmt.Print()
			fmt.Printf("v1 type:%T\n", v.(map[string]interface{})["id"])
			if v.(map[string]interface{})["id"] == id {
				data = append(data[:k], data[k+1:]...)
				break
			}
		}
		Map["Rows"] = data
		utils.WriteJson("conf/ButtonConf.json", Map)

	}
}
