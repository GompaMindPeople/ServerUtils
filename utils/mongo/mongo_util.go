/*
@Time : 2019/6/10 13:56 
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package mongo

import "gopkg.in/mgo.v2"


type MGInterface interface {
	GetConnection()
}

type MG struct {
	HostName string
	session *mgo.Session
}



/**
	获取一个session
 */
func (mg *MG) GetSession() *mgo.Session {
	session, err := mgo.Dial(mg.HostName)
	if err != nil {
		panic(err)
	}
	mg.session = session
	return session
}



