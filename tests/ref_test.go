/*
@Time : 2019/6/24 17:08
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package test

import (
	"ServerUtils/models"
	"reflect"
	"testing"
)

func TestRef(t *testing.T) {
	table := models.TypeTable{}
	valueOf := reflect.ValueOf(&table)
	elem := valueOf.Elem()
	name := elem.FieldByName("TypeName")
	name.SetString("sdad")

	t.Log(table.TypeName)
}
