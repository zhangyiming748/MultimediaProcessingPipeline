package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	mylog "Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/model"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/sql"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	DP "github.com/OwO-Network/DeepLX/translate"
)

// go test -timeout 2000m -v -run TestTransAll
// docker run -dit --name trans -v /c/Users/zen/Github/MultimediaProcessingPipeline:/app -v /c/Users/zen/Videos/export/sdde:/data zhangyiming748/stand:latest bash
func TestTransAll(t *testing.T) {
	defer func() {
		log.Println("全部任务完成")
	}()

	p := &constant.Param{
		Root:     "E:\\Downloads\\My Pack\\anime\\动画\\蒂法 三部",
		Language: "English",
		Pattern:  "mp4",
		Model:    "medium.en",
		Location: "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline",
		Proxy:    "192.168.1.35:8889",
	}

	mylog.SetLog(p)
	sql.SetLevelDB(p)
	sql.SetMysql()
	sql.GetMysql().Sync2(model.TranslateHistory{})
	replace.SetSensitive(p)
	//util.ExitAfterRun()
	fps := getFiles(p.GetRoot())

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

func TestDeeplx(t *testing.T) {
	src := "hЯ суперкрасивый парень.ello"
	token := ""
	os.Setenv("TOKEN", token)
	deeplx, err := Deeplx(src)
	if err != nil {
		t.Log(err)
	}
	t.Log(deeplx)
}
func TestTranslateByDeepLX(t *testing.T) {
	source := "auto"
	target := "zh"
	text := "hello"
	lx, err := DP.TranslateByDeepLX(source, target, text, "html", "", "")
	if err != nil {
		log.Fatalln(err)
	}
	t.Logf("%+v\ndata = %v\n", lx, lx.Data)
}
func TestTransByDeeplx(t *testing.T) {
	ch := make(chan string, 1)
	var one sync.Once
	var wg sync.WaitGroup
	wg.Add(1)
	//os.Setenv("TOKEN", "")
	go TransByDeeplx("hello", "http://127.0.0.1:8889", &one, &wg, ch)
	wg.Wait()
	fmt.Println("通道接收 ", <-ch)
}

// go test -timeout 2000m -v -run TestSingleTranslate
func TestSingleTranslate(t *testing.T) {
	p := new(constant.Param)
	p.SetProxy("http://127.0.0.1:8889")
	c := new(constant.Count)
	beforeWords := util.ReadByLine("before.txt")
	afterWords := []string{}
	for i, word := range beforeWords {
		log.Printf("第%d个句子:%s\n", i, word)
		after := Translate(word, p, c)
		after = strings.Replace(after, "\n", "", 1)
		afterWords = append(afterWords, after)
	}
	util.WriteByLineOnce("after.txt", afterWords)
}
