/*
@Time : 2019/5/24 17:25
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package test

import (
	"ServerUtils/server"
	"bytes"
	"fmt"
	"testing"
)

func TestRunSsh(t *testing.T) {
	//bean := server.SshBean{}
	//cmd:= "cd  /home/007/server/game-server2 &&  tail -1000 game-server.log   "
	//var	stdout ,stderr  bytes.Buffer
	//bean.RunSsh(cmd,&stdout,&stderr)
	//logs.Debug(string(stdout.Bytes()))
	//bean.SendShell("echo 588888888888888;")
}

func TestSendMessage(t *testing.T) {
	session, e := server.GetSSHSession()
	if e == nil {
		cmd := "tail -1000  /home/007/server/game-server2/game-server.log"
		var stdOut, stdErr bytes.Buffer
		session.SendMessage(cmd, &stdOut, &stdErr)
		fmt.Print(stdOut)
	}

	//cmd1:= "./server.check.sh start;"
	//var stdOut,stdErr bytes.Buffer
	//session.SendMessage(cmd1,&stdOut,&stdErr)

	if e != nil {
		fmt.Print(e)
	}
	for true {
		fmt.Print()
	}
}

func TestGetSSHConnect(t *testing.T) {
	//var stdOut,stdErr bytes.Buffer
	//cmd := "cd /home/007/server/game-server2/_shell;./server.check.sh stop;"
	//cmdlist := strings.Split(cmd, ";")
	//if session,_, e := server.getSSHConnect("root", "2098231.xzm", "192.168.5.189", 22);e == nil{
	//	session.Stdout = &stdOut
	//	session.Stderr = &stdErr
	//	closer, e := session.StdinPipe()
	//	defer session.Close()
	//	err := session.Shell()
	//	if err != nil{
	//		fmt.Println("错误err-->",e)
	//	}
	//	for _, c := range cmdlist {
	//		c = c + "\n"
	//		n, e := closer.Write([]byte(c))
	//		if e != nil{
	//			fmt.Println("错误",e)
	//		}else{
	//			fmt.Println("--->",n)
	//		}
	//	}
	//	//e = session.Run("")
	//	t.Log(stdOut.String())
	//	if e != nil{
	//		fmt.Println("错误",e)
	//	}
	//	//e = session.Wait()
	//	t.Log(stdOut.String() + stdErr.String())
	//
	//}else{
	//	log.Fatal("error->",e)
	//}

}
