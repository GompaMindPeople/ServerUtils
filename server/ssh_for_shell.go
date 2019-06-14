/*
@Time : 2019/5/24 15:56
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package server

import (
	"ServerUtils/models"
	"ServerUtils/utils"
	"bytes"
	"fmt"
	"github.com/astaxie/beego/logs"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

type SSHModel struct {
	sshService *ssh.Session
	closer     io.WriteCloser
	out        io.Reader
}

var sshModel *SSHModel

var mutex sync.Mutex

type SshBean struct {
	Client  *ssh.Client
	Session *ssh.Session
}

//获取ssh客户端实例
func (s *SshBean) GetClient(user, password, host string, port int16) (*ssh.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	s.Client = client
	return client, nil
}

//获取session,始终获得最新的实例
func (s *SshBean) GetSession() (*ssh.Session, error) {
	client := s.Client
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SshBean) RunSsh(cmd string, id int, stdout *bytes.Buffer, stderr *bytes.Buffer) error {
	config := models.SshConfig{}
	config.SelectOne(id)
	client, err := s.GetClient(config.UserName, config.Password, config.HostName, config.Port)
	if err != nil {
		logs.Error(err)
		return err
	}
	session, err := client.NewSession()
	//stdinBuf, err := session.StdinPipe()
	if err != nil {
		logs.Error(err)
	}
	//cmdlist := strings.Split(cmd, ";")
	defer session.Close()
	session.Stdout = stdout
	session.Stderr = stderr
	//err = session.Run(cmd)
	//err = session.Shell()
	err = session.Start(cmd)
	if err != nil {
		logs.Error(err)
		return err
	}
	//for _, c := range cmdlist {
	//   c = c + "\n"
	//   stdinBuf.Write([]byte(c))
	//}
	err = session.Wait()
	//logs.Error(stdout.String() + stderr.String())
	//err = session.Run("echo 123456")
	//ret, err := strconv.Atoi( str.Replace( stdOut.String(), "\n", "", -1 )  )
	if err != nil {
		return err
	}

	//	fmt.Printf("%d, %s\n", ret, stdErr.String() )
	return nil
}

// RFC 4254 Section 6.5.
type execMsg struct {
	Command string
}

func (s *SshBean) SendShell(cmd string) {
	req := execMsg{
		Command: cmd,
	}

	//ok, err := s.ch.SendRequest("exec", true, Marshal(&req))
	b, e := s.Session.SendRequest("exec", true, ssh.Marshal(&req))
	if e != nil {
		logs.Error(e)
	}
	log.Print(b)
}

/**
获取一个 封装后的 SSHSession 会话
单例模式,只获得一个session 会话
*/
func GetSSHSession() (*SSHModel, error) {
	mutex.Lock()
	defer mutex.Unlock()
	var e error
	if sshModel == nil {
		config := utils.GetSSHConfig()
		if &config == nil {
			log.Fatal("")
		}
		if session, closer, out, e := getSSHConnect(config.UserName, config.Password, config.HostName, config.Port); e == nil {
			sshModel = &SSHModel{session, closer, out}
			return sshModel, nil
		}
		return nil, e
	}

	return sshModel, nil
}

func (model *SSHModel) Close() {
	err := model.closer.Close()
	model.sshService.Close()
	sshModel = nil
	if err != nil {
		log.Fatal("关闭closer发生错误-->", err)
	}
}

func (model *SSHModel) SendMessage(cmd string, stdOut, stdErr *bytes.Buffer) {

	//var stdOut1 []byte
	cmdList := strings.Split(cmd, ";")
	//session := model.sshService
	//defer model.Close()
	//session.Stdout = stdOut
	//session.Stderr = stdErr
	//closer:= model.closer
	for _, c := range cmdList {
		//c = c + "\n"
		//_, e := closer.Write([]byte(c))
		////output, e := session.CombinedOutput(c)
		//if e != nil{
		//	fmt.Println("写出ssh消息时发生错误-->",e)
		//}
		model.sshService.Run(c)
	}
	//_, e := model.out.Read(stdOut1)
	//fmt.Println("-->",stdOut1)
	//fmt.Println("-->",e)
	//all, e := ioutil.ReadAll(model.out)
	//if e != nil{
	//	logs.Error(e)
	//}
	//fmt.Print(all)
}

//func GetSSHConnect(){
//
//}

func getSSHConnect(user, password, host string, port int16) (*ssh.Session, io.WriteCloser, io.Reader, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
		closer       io.WriteCloser
		out          io.Reader
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, nil, nil, err
	}
	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, nil, nil, err
	}
	if closer, err = session.StdinPipe(); err != nil {
		return nil, nil, nil, err
	}
	if out, err = session.StdoutPipe(); err != nil {
		return nil, nil, nil, err
	}
	err = session.Shell()
	if err != nil {
		fmt.Println("打开shell发生错误-->", err)
	}
	return session, closer, out, nil
}
