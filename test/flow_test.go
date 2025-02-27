package t

import (
	"Multimedia_Processing_Pipeline/constant"
	mylog "Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/model"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/sql"
	trans "Multimedia_Processing_Pipeline/translate"
	"Multimedia_Processing_Pipeline/util"
	"Multimedia_Processing_Pipeline/whisper"
	"Multimedia_Processing_Pipeline/ytdlp"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var p = &constant.Param{
	VideosLocation: "/Videos",
	Language:       "English",
	Pattern:        "mp4",
	Model:          "medium.en",
	ToolsLocation:  "/app/test",
	Proxy:          "http://192.168.2.8:8889",
	Merge:          false,
	//Lines:          string // 保存下载url的文档 默认放在root下 文件名为 link.list
	MysqlUser:    "root",
	MysqlPass:    "163453",
	MysqlHost:    "192.168.2.5",
	MysqlPort:    "3306",
	TransService: "192.168.2.5:8192",
}

func init() {
	mylog.SetLog(p)
	sql.SetLevelDB(p)
	sql.SetMysql()
	sql.GetMysql().Sync2(model.TranslateHistory{})
	replace.SetSensitive(p)
}

// go test -timeout 2000h -v -run TestYTdlp
func TestYTdlp(t *testing.T) {
	file, err := os.OpenFile("fail.list", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	link := filepath.Join(p.GetToolsLocation(), "link.list")
	log.Println("开始读文件")
	uris := util.ReadByLine(link)
	for _, uri := range uris {
		if uri == "" {
			continue
		}
		if strings.HasPrefix(uri, "#") {
			continue
		}
		if link := ytdlp.DownloadVideo(uri, p); link == "" {
			file.WriteString(fmt.Sprintln(uri))
		}
	}
	file.Sync()
}

// go test -timeout 2000h -v -run TestWhisper
func TestWhisper(t *testing.T) {

	fps := getFiles(p.GetVideosLocation())
	cmds := []string{}
	for _, fp := range fps {
		if strings.HasSuffix(fp, p.GetPattern()) {
			cmd := whisper.GetSubtitle(fp, p, false)
			cmds = append(cmds, cmd)
		}
	}
	if runtime.GOOS == "windows" {
		fp := filepath.Join(p.GetVideosLocation(), "whisper.ps1")
		util.WriteByLine(fp, cmds)
	} else {
		fp := filepath.Join(p.GetVideosLocation(), "whisper.sh")
		util.WriteByLine(fp, cmds)
	}
}

// go test -timeout 2000h -v -run TestTransAll
func TestTransAll(t *testing.T) {
	//util.ExitAfterRun()
	fps := getFiles(p.GetVideosLocation())
	log.Println(fps)
	c := new(constant.Count)
	for _, fp := range fps {
		if strings.HasSuffix(fp, ".srt") {
			trans.Trans(fp, p, c)
		}
	}
}

// go test -timeout 2000h -v -run TestTransFileAndArchive
func TestTransFileAndArchive(t *testing.T) {
	//util.ExitAfterRun()
	fps := getFiles(p.GetVideosLocation())
	log.Println(fps)
	for _, fp := range fps {
		if strings.HasSuffix(fp, ".txt") {
			trans.TransFile(fp, p)
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
