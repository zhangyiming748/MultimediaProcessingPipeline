package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	mylog "Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// go test -timeout 2000m -v -run TestTransAll

func TestTransAll(t *testing.T) {
	p := &constant.Param{
		Root:     "/data",
		Language: "English",
		Pattern:  "mp4",
		Model:    "base",
		Location: "/data",
		Proxy:    "192.168.1.20:8889",
	}
	mylog.SetLog(p)
	sql.SetDatabase(p)
	//util.ExitAfterRun()
	fps := getFiles(p.GetRoot())

	c := new(constant.Count)
	for _, fp := range fps {
		if strings.HasSuffix(fp, ".mp4") {
			Trans(fp, p, c)
		}
	}
}

func getFiles(currentDir string) (filePaths []string) {
	// 获取当前工作目录
	//currentDir, err := os.Getwd()
	//if err != nil {
	//	fmt.Println("获取当前目录失败:", err)
	//	return []string{}
	//}

	// 创建一个切片来保存文件的绝对路径
	//var filePaths []string

	// 使用 Walk 函数遍历当前目录
	err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查是否是文件
		if !info.IsDir() {
			filePaths = append(filePaths, path) // 将文件的绝对路径添加到切片
		}
		return nil
	})

	if err != nil {
		fmt.Println("遍历目录失败:", err)
		return
	}

	// 打印所有文件的绝对路径
	for _, filePath := range filePaths {
		fmt.Println(filePath)
	}
	return filePaths
}
