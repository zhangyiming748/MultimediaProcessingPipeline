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
	"path/filepath"
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
	p := new(constant.Param)
	p.Root = "/home/zen/git/MultimediaProcessingPipeline/ytdlp"
	p.Language = "English"
	p.Pattern = "mp4"
	p.Model = "base"
	p.Location = "/home/zen/git/MultimediaProcessingPipeline/ytdlp"
	p.Proxy = "192.168.1.20:8889"
	p.Merge = false

	mylog.SetLog(p)
	sql.SetLevelDB(p)
	//util.ExitAfterRun()
	replace.SetSensitive(p)
	c := new(constant.Count)
	fps := []string{
		"/home/zen/git/MultimediaProcessingPipeline/ytdlp/Anastasia Doll： Bubbles, Boobs & Beyond.mp4",
		"/home/zen/git/MultimediaProcessingPipeline/ytdlp/CATALINA CRUZ - Huge Dildo Plunges Deep Into Wet Pussy.mp4",
		"/home/zen/git/MultimediaProcessingPipeline/ytdlp/Danielle Derek： Hot Stack.mp4",
		"/home/zen/git/MultimediaProcessingPipeline/ytdlp/Huge tits Ricki Raxxx bra changing show .mp4",
		"/home/zen/git/MultimediaProcessingPipeline/ytdlp/Tanya Virago Cums Clean.mp4",
		"/home/zen/git/MultimediaProcessingPipeline/ytdlp/Tanya Virago： Huge Tits, Gaping Pussy.mp4",
	}
	for _, fp := range fps {
		Trans(fp, p, c)
	}
}

// go test -v -run TestTransJapanese
func TestTransJapanese(t *testing.T) {
	p := new(constant.Param)
	p.Root = "/mnt/c/Users/zen/Downloads"
	p.Language = "Japanese"
	p.Pattern = "mp4"
	p.Model = "small"
	p.Location = "/mnt/c/Users/zen/Downloads"
	p.Proxy = "192.168.1.20:8889"
	p.Merge = false
	mylog.SetLog(p)
	sql.SetLevelDB(p)
	//util.ExitAfterRun()
	replace.SetSensitive(p)
	c := new(constant.Count)
	Trans("/mnt/c/Users/zen/Github/Multimedia_Processing_Pipeline/Tom Cruise Terrifies James in 'Top Gun' Fighter Jet! [v1iZtBM23bY].srt", p, c)
}
func TestSplitExt(t *testing.T) {
	name := "1111.cap"
	ext := filepath.Ext(name)
	ext2 := filepath.Ext(name)
	t.Log(ext, ext2)
}
