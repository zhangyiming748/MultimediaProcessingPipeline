package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/sql"
	"fmt"
	"github.com/zhangyiming748/DeepLX"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// go test -v -run TestTranslate
//func TestDeepXl(t *testing.T) {
//	ret := DeepXl("hello,world")
//	t.Log(ret)
//}

func TestMaster(t *testing.T) {

	url := "http://192.168.1.6:1188/translate"
	method := "POST"

	payload := strings.NewReader(`{
    "text": "Hello, world!",
    "source_lang": "auto",
    "target_lang": "ZH"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func TestDeepLX(t *testing.T) {
	lx, err := DeepLx.TranslateByDeepLX("auto", "zh", "hello world", "")
	if err != nil {
		return
	} else {
		t.Log(lx)
	}
}

// go test -v -run TestWhisper
func TestTrans(t *testing.T) {
	p := &constant.Param{
		Root:     "/home/zen/git/MultimediaProcessingPipeline/ytdlp",
		Language: "English",
		Pattern:  "mp4",
		Model:    "base",
		Location: "/home/zen/git/MultimediaProcessingPipeline/ytdlp",
		Proxy:    "192.168.1.20:8889",
	}
	sql.SetDatabase(p)
	log.SetLog(p)
	c := new(constant.Count)
	Trans("/home/zen/git/MultimediaProcessingPipeline/ytdlp/EDGE FOR ME -JOI [656dc0089a6eb].mp4", p, c)
}
