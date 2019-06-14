/*
@Time : 2019/5/24 16:06
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package models

import (
	"ServerUtils/Global"
)

type SshConfig struct {
	Id       int    `orm:"column(id);auto" description:"id"`
	UserName string `orm:"column(UserName)"`
	Password string `orm:"column(Password)"`
	HostName string `orm:"column(HostName)"`
	Port     int16  `orm:"column(Port)"`
}

/**
插入一个对象
*/
func (config *SshConfig) InsertOne() (int64, error) {
	id, err := Global.OrmInstance.Insert(config)
	return id, err
}

/**
更新数据
*/
func (config *SshConfig) Update() (int64, error) {
	i, e := Global.OrmInstance.Update(config)
	return i, e
}

/**
查询 sshConfig
*/
func (config *SshConfig) SelectOne(Id int) *SshConfig {
	config.Id = Id
	err := Global.OrmInstance.Read(config)
	if err != nil {
		panic(err)
	}
	return config
}
