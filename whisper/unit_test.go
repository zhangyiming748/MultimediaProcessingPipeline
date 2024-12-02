package whisper

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	//"time"
)

// go test -timeout 2000h -v -run TestWhisper
func TestWhisper(t *testing.T) {
	p := &constant.Param{
		Root:     "E:\\Downloads\\My Pack\\anime\\动画\\蒂法 三部",
		Language: "English",
		Pattern:  "mp4",
		Model:    "large-v3",
		Location: "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline",
		Proxy:    "127.0.01:8889",
	}
	log.SetLog(p)
	fps := getFiles(p.GetRoot())
	cmds := []string{}
	for _, fp := range fps {
		if strings.HasSuffix(fp, p.GetPattern()) {
			cmd := GetSubtitle(fp, p, true)
			cmds = append(cmds, cmd)
		}
	}
	if runtime.GOOS == "windows" {
		fp := filepath.Join(p.GetRoot(), "whisper.bat")
		util.WriteByLine(fp, cmds)
	} else {
		fp := filepath.Join(p.GetRoot(), "whisper.sh")
		util.WriteByLine(fp, cmds)
	}
	util.WriteByLine("whisper.bat", cmds)
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
			GetSubtitle(fp, p, false)
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
