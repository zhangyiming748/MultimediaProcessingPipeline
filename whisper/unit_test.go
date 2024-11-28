package whisper

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	//"time"
)

// go test -timeout 2000h -v -run TestWhisper
func TestWhisper(t *testing.T) {
	p := &constant.Param{
		Root:     "/data/jp",
		Language: "Japanese",
		Pattern:  "mp4",
		Model:    "large-v3",
		Location: "/data",
		Proxy:    "192.168.1.31:8889",
	}
	log.SetLog(p)
	fps := getFiles(p.GetRoot())
	for _, fp := range fps {
		if strings.HasSuffix(fp, p.GetPattern()) {
			GetSubtitle(fp, p)
		}
	}
}
func TestWhisperOnWindows(t *testing.T) {
	p := &constant.Param{
		Root:     "C:\\Users\\zen\\Ada Hunter Red Suit All Cutscenes",
		Language: "English",
		Pattern:  "m4a",
		Model:    "large-v3",
		Location: "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline",
		Proxy:    "192.168.1.31:8889",
	}
	log.SetLog(p)
	fps := getFiles(p.GetRoot())
	for _, fp := range fps {
		if strings.HasSuffix(fp, p.GetPattern()) {
			GetSubtitle(fp, p)
		}
	}
}

func getFiles(currentDir string) (filePaths []string) {
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

// go test -timeout 2000h -v -run TestGetSpecif

func TestGetSpecif(t *testing.T) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	t.Log(hostname)
}
