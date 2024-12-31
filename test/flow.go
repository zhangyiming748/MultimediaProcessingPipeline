package t

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/util"
	"Multimedia_Processing_Pipeline/whisper"
	"Multimedia_Processing_Pipeline/ytdlp"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var p = &constant.Param{
	Root:     "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline\\ytdlp",
	Language: "English",
	Pattern:  "mp4",
	Model:    "small",
	Location: "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline\\ytdlp",
	Proxy:    "192.168.1.35:8889",
}

func TestYTdlp(t *testing.T) {

	file, err := os.OpenFile("fail.list", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	log.SetLog(p)
	link := filepath.Join(p.GetLocation(), "link.list")
	uris := util.ReadByLine(link)
	for _, uri := range uris {
		if link := ytdlp.DownloadVideo(uri, p); link == "" {
			file.WriteString(fmt.Sprintln(uri))
		}
	}
	file.Sync()
}

// go test -timeout 2000h -v -run TestWhisper
func TestWhisper(t *testing.T) {
	p := &constant.Param{
		Root:     "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline\\videos",
		Language: "English",
		Pattern:  "mp4",
		Model:    "medium.en",
		Location: "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline",
		Proxy:    "192.168.1.35:8889",
	}
	log.SetLog(p)
	fps := getFiles(p.GetRoot())
	cmds := []string{}
	for _, fp := range fps {
		if strings.HasSuffix(fp, p.GetPattern()) {
			cmd := whisper.GetSubtitle(fp, p, false)
			cmds = append(cmds, cmd)
		}
	}
	if runtime.GOOS == "windows" {
		fp := filepath.Join(p.GetRoot(), "whisper.ps1")
		util.WriteByLine(fp, cmds)
	} else {
		fp := filepath.Join(p.GetRoot(), "whisper.sh")
		util.WriteByLine(fp, cmds)
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
