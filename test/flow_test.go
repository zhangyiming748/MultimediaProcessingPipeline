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
	VideosLocation: "/Users/zen/github/MultimediaProcessingPipeline/Zoey Holloway/test",
	Language:       "English",
	Pattern:        "mp4",
	Model:          "large-v3",
	ToolsLocation:  "/Users/zen/github/MultimediaProcessingPipeline",
	Proxy:          "http://192.168.2.10:8889",
	Merge:          false,
	//Lines:          string // 保存下载url的文档 默认放在root下 文件名为 link.list
	MysqlUser:    "root",
	MysqlPass:    "163453",
	MysqlHost:    "192.168.2.10",
	MysqlPort:    "3306",
	TransService: "http://192.168.2.10:8192",
}

func init() {
	// 初始化全局变量 p
	mylog.SetLog(p)
	sql.SetMysql(p)
	sql.GetMysql().Sync2(model.TranslateHistory{})
	sql.GetMysql().Sync2(model.Sensitive{})

	// 优先从环境变量获取 apikey
	if apikey := os.Getenv("DeepLX"); apikey != "" {
		p.LinuxDo = apikey
		log.Printf("从环境变量获取到 DeepLX: %v", p.LinuxDo)
	} else {
		// 如果环境变量不存在，则尝试从文件读取
		readKey(p)
	}

	log.SetFlags(log.Ltime | log.Lshortfile)
	replace.SetSensitive(p)
}

func readKey(p *constant.Param) {
	if _, err := os.Stat("apikey"); os.IsNotExist(err) {
		log.Println("apikey文件不存在")
		return
	}
	keys := util.ReadByLine("apikey")
	p.LinuxDo = keys[0]
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
		if link := ytdlp.DownloadVideo(uri, p.GetProxy(), p.GetVideosLocation()); link == "" {
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
			cmd := whisper.GetSubtitle(fp, p.GetModel(), p.GetToolsLocation(), p.GetLanguage(), p.GetVideosLocation())
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
	// 检查目录是否存在
	if _, err := os.Stat(p.VideosLocation); os.IsNotExist(err) {
		t.Fatal("目录不存在:", err)
	}

	fps := getFiles(p.VideosLocation)
	log.Printf("找到 %d 个文件", len(fps))
	if len(fps) == 0 {
		t.Fatal("未找到任何文件")
	}
	srtCount := 0
	for _, fp := range fps {
		if strings.HasSuffix(fp, ".srt") {
			log.Printf("处理文件: %s", fp)
			srtCount++
			trans.Trans(fp)

		}
	}
	if srtCount == 0 {
		t.Fatal("未找到任何.srt文件")
	}
}
func TestCache(t *testing.T) {
	c := new(model.TranslateHistory)
	c.Src = "hello"
	has, err := c.FindBySrc()
	if err != nil {
		t.Fatal(err)
	} else if has {
		t.Log(c.Dst)
	} else {
		t.Log("not found")
	}
}

// go test -timeout 2000h -v -run TestTransFileAndArchive

// go test -v -run TestGetEnv
func TestGetEnv(t *testing.T) {
	// 打印所有环境变量
	t.Log("=== 开始打印所有环境变量 ===")
	envs := os.Environ()
	t.Logf("总共有 %d 个环境变量", len(envs))
	for _, env := range envs {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			t.Logf("变量名: %s\n值: %s", parts[0], parts[1])
		}
	}
	t.Log("=== 环境变量打印完成 ===")
	path := os.Getenv("DeepLX")
	if path == "" {
		t.Error("无法获取 DeepLX 环境变量")
	} else {
		t.Log(path)
	}

	// 测试获取自定义环境变量
	testKey := "TEST_ENV_VAR"
	testValue := "test_value"
	os.Setenv(testKey, testValue)

	value := os.Getenv(testKey)
	if value != testValue {
		t.Errorf("期望获取的环境变量值为 %s，实际获取到的是 %s", testValue, value)
	}

	// 测试获取不存在的环境变量
	nonExistValue := os.Getenv("NON_EXIST_ENV_VAR")
	if nonExistValue != "" {
		t.Error("获取不存在的环境变量应该返回空字符串")
	}
}

func getFiles(currentDir string) (filePaths []string) {
	err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查是否是文件
		if !info.IsDir() {
			if !strings.HasPrefix(info.Name(), ".") {
				filePaths = append(filePaths, path) // 将文件的绝对路径添加到切片
			}
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
