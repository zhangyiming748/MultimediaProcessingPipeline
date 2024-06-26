package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
执行命令过程中可以循环打印消息
*/
func ExecCommand(c *exec.Cmd, msg string) (e error) {
	log.Printf("开始执行命令:%v\n", c.String())
	stdout, err := c.StdoutPipe()
	c.Stderr = c.Stdout
	if err != nil {
		log.Printf("连接Stdout产生错误:%v\n", err)
		return err
	}
	if err = c.Start(); err != nil {
		log.Printf("启动cmd命令产生错误:%V\n", err)
		return err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = strings.Replace(t, "\u0000", "", -1)
		fmt.Printf("\r%v\n%v\n", t, msg)
		if err != nil {
			break
		}
	}
	if err = c.Wait(); err != nil {
		log.Printf("命令执行中产生错误:%v\n", err)
		return err
	}
	if isExitLabel() {
		log.Fatalf("命令端获取到退出状态,命令结束后退出:%v\n", c.String())
	}
	if GetExitStatus() {
		log.Fatalf("命令端获取到退出状态,命令结束后退出:%v\n", c.String())
	}
	return nil
}
func ExecCommand4YtdlpDestination(c *exec.Cmd, msg string) (Destination string, e error) {
	log.Printf("开始执行命令:%v\n", c.String())

	stdout, err := c.StdoutPipe()
	c.Stderr = c.Stdout
	if err != nil {
		log.Printf("连接Stdout产生错误:%v\n", err)
		return "", err
	}
	if err = c.Start(); err != nil {
		log.Printf("启动cmd命令产生错误:%V\n", err)
		return "", err
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		t := string(tmp)
		t = strings.Replace(t, "\u0000", "", -1)
		fmt.Printf("\r%v\n%v\n", t, msg)
		if strings.Contains(t, "Destination") {
			strings.Split(t, " Destination: ")
		}
		if err != nil {
			break
		}
	}
	if err = c.Wait(); err != nil {
		log.Printf("命令执行中产生错误:%v\n", err)
		return "", err
	}
	if isExitLabel() {
		log.Fatalf("命令端获取到退出状态,命令结束后退出:%v\n", c.String())
	}
	if GetExitStatus() {
		log.Fatalf("命令端获取到退出状态,命令结束后退出:%v\n", c.String())
	}
	return Destination, nil
}

/*
判断古希腊掌管退出信号的文件是否存在
*/
func isExitLabel() bool {
	filePath := "/exit"
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("古希腊掌管退出信号的文件不存在")
		return false
	} else {
		fmt.Println("古希腊掌管退出信号的文件存在")
		return true
	}
}
