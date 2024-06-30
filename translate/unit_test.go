package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	mylog "Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/replace"
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
		Root:     "/Users/zen/Downloads",
		Language: "Japanese",
		Pattern:  "mp3",
		Model:    "base",
		Location: "/Users/zen/Downloads",
		Proxy:    "192.168.1.20:8889",
	}
	sql.SetDatabase(p)
	mylog.SetLog(p)
	c := new(constant.Count)
	fps := []string{
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD1/01 - アーニャとベッキーのびっくり大作戦！.mp3",
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD1/02 - アーニャとベッキーのびっくり大作戦！ Cast Commentary.mp3",
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD2/01 - ブライア姉弟のスペシャルクッキング.mp3",
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD2/02 - ブライア姉弟のスペシャルクッキング Cast Commentary.mp3",
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD3/01 - ケーキを選んで世界平和⁉.mp3",
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD3/02 - ケーキを選んで世界平和⁉ Cast Commentary.mp3",
	}
	for _, fp := range fps {
		Trans(fp, p, c)
	}
}

// go test -v -run TestTransJapanese
func TestTransJapanese(t *testing.T) {
	p := new(constant.Param)
	p.Root = "/Users/zen/Github/MultimediaProcessingPipeline/ytdlp"
	p.Language = "Japanese"
	p.Pattern = "mp4"
	p.Model = "base"
	p.Location = "/Users/zen/Github/MultimediaProcessingPipeline/ytdlp"
	p.Proxy = "192.168.1.20:8889"
	p.Merge = false

	mylog.SetLog(p)
	sql.SetDatabase(p)
	//util.ExitAfterRun()
	replace.SetSensitive(p)

	c := new(constant.Count)
	Trans("/Users/zen/Github/MultimediaProcessingPipeline/test/NieR：Automata Fan Festival 12022 koncert [wX_SAi_ZcFQ].mp4", p, c)
}
