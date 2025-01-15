package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	mylog "Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/model"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// go test -timeout 2000m -v -run TestTransAll
// docker run -dit --name trans -v /c/Users/zen/Github/MultimediaProcessingPipeline:/app -v /c/Users/zen/Videos/export/sdde:/data zhangyiming748/stand:latest bash
func TestTransAll(t *testing.T) {
	defer func() {
		log.Println("全部任务完成")
	}()

	p := &constant.Param{
		VideosLocation: "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline\\videos\\2",
		Language:       "English",
		Pattern:        "mp4",
		Model:          "medium.en",
		ToolsLocation:  "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline\\videos\\2",
		Proxy:          "192.168.1.35:8889",
		Mysql:          "192.168.1.9:3306",
		TransService:   "192.168.1.9:3389",
	}

	mylog.SetLog(p)
	sql.SetLevelDB(p)
	sql.SetMysql(p)
	sql.GetMysql().Sync2(model.TranslateHistory{})
	replace.SetSensitive(p)
	//util.ExitAfterRun()
	fps := getFiles(p.GetVideosLocation())

	c := new(constant.Count)
	for _, fp := range fps {
		if strings.HasSuffix(fp, ".srt") {
			Trans(fp, p, c)
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
