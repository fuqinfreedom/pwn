package main

import (
	"bytes"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
  result()
}
//整体调用
func result()  {
	exec.Command("echo '' > /home/kali/RiskDetector/account/account.txt")
	exec.Command("echo '' > /home/kali/RiskDetector/account/loguser.txt")
	exec.Command("echo '' > /home/kali/RiskDetector/account/passwd.txt")
	exec.Command("cat /home/kali/RiskDetector/account/loguser.txt")
	fmt.Println("============")

	re,err := sshL("47.105.149.182","root","fuqinsec@123","password","","cat /etc/passwd ",22)
	print(string(re))
	re2,err := sshL("47.105.149.182","root","fuqinsec@123","password","","echo $SHELL",22)
	print(string(re2))
	check(err)
	write("/home/kali/RiskDetector/account/passwd.txt",string(re))
	read("/home/kali/RiskDetector/account/passwd.txt","/home/kali/RiskDetector/account/account.txt",string(re2))
	grep("/home/kali/RiskDetector/account/account.txt","/home/kali/RiskDetector/account/loguser.txt")

}
//获取passwd文件内容
func sshL(sshHost,sshUser,sshPassword,sshType,sshKeyPath,command string,sshPort int) ([]byte,error) {

	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second,//ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	}
	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("创建ssh client 失败",err)
		return nil,err
	}
	defer sshClient.Close()


	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session 失败",err)
		return nil,err
	}
	defer session.Close()
	//执行远程命令
	combo,err := session.CombinedOutput(command)
	if err != nil {
		log.Fatal("远程执行cmd 失败",err)
		return nil,err
	}
	return combo,nil
}
//配合获取passwd文件内容
func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
//passwd文件内容写入
func write(filename string ,txt string)  {
	dstFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer dstFile.Close()
	s := string(txt)
	dstFile.WriteString(s + "\n")
	dstFile.Close()


}
//将passwd以每行结尾是否使用shell进行过滤，主要功能在read脚本中
func read(stringset ...string) {

	cmd, _ := exec.Command("/home/kali/RiskDetector/account/read.sh", stringset...).Output()
	// cmd,_:=exec.Command("ping","127.0.0.1" ,"-n","1").Output()
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(cmd), simplifiedchinese.GB18030.NewDecoder()))
	fmt.Println(data)
}
//提取出可登陆的账户名
func grep(stringset ...string) {

	cmd, _ := exec.Command("/home/kali/RiskDetector/account/grep.sh", stringset...).Output()
	// cmd,_:=exec.Command("ping","127.0.0.1" ,"-n","1").Output()
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader(cmd), simplifiedchinese.GB18030.NewDecoder()))
	fmt.Println(data)
}
//错误检查
func check(e error) {
	if e != nil {
		panic(e)
	}
}