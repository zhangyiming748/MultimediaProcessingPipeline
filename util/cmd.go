package util

import (
	"Multimedia_Processing_Pipeline/replace"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
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
			sp := strings.Split(t, " Destination: ")
			Destination = sp[1]
			log.Printf("Destination捕获到视频标题:%s", Destination)
		} else if strings.Contains(t, "container of") {
			content, _ := getQuotedContent(t)
			Destination = content
			log.Printf("container of捕获到视频标题:%s", Destination)
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
	RealName := replace.RealName(Destination)
	return RealName, nil
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

/*
获取引号部分中的内容
*/
func getQuotedContent(input string) (title string, err error) {
	re := regexp.MustCompile(`"(.*?)"`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		log.Printf("捕获到引号中可能是视频标题的内容:%s\n", matches[1])
		return matches[1], nil
	} else {
		return "", errors.New("未找到匹配项")
	}
}
