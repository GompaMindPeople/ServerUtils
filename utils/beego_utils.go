/*
@Time : 2019/6/13 11:03
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package utils

import (
	"log"
	"reflect"
)

func FillObjectByMap(obj interface{}, Map interface{}) interface{} {
	//elem := reflect.ValueOf(&obj).Elem()
	//names := GetFieldName(&obj)

	//for  k,v :=  range names  {
	//	elem.FieldByName(v).SetInt(Map[v])
	//}
	return nil

}

func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}
