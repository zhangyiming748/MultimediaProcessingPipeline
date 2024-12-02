package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ReadByLine(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		log.Println("按行读文件出错")
		return []string{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines, string(a))
	}
	return lines
}

// 按行写文件
func WriteByLine(fp string, s []string) {
	file, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range s {
		writer.WriteString(v)
		writer.WriteString("\n")
	}
	writer.Flush()
}

// 按行写文件 截断
func WriteByLineOnce(fp string, s []string) {
	file, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range s {
		writer.WriteString(v)
		writer.WriteString("\n")
	}
	writer.Flush()
}
func IsExist(folderPath string) bool {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		log.Printf("文件夹:%v不存在\n", folderPath)
		return false
	} else {
		log.Printf("文件夹:%v存在\n", folderPath)
		return true
	}
}
func IsExistCmd(cmds ...string) bool {
	for _, cmd := range cmds {
		//cmd := "ls" // 需要测试的命令
		_, err := exec.LookPath(cmd)
		if err != nil {
			log.Printf("命令:%s不存在\n", cmd)
			return false
		} else {
			log.Printf("命令:%s存在\n", cmd)
		}
	}
	return true
}
func GetAllFileInfoFast(dir, pattern string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), pattern) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}
func ReadInSlice(fp string) []string {
	fileBytes, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []string{}
	}
	// 创建一个bufio.Reader对象
	reader := bufio.NewReader(bytes.NewReader(fileBytes))
	// 按行读取文件内容并存储到字符串切片中
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	return lines
}
