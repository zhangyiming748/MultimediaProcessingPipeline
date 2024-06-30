package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/sql"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"github.com/zhangyiming748/DeepLX"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	seed = rand.New(rand.NewSource(time.Now().Unix()))
)

const (
	TIMEOUT = 8 //second
	RETRY   = 3
)

func Translate(src string, p *constant.Param, c *constant.Count) string {
	//trans -brief ja:zh "私の手の動きに合わせて|そう"
	var dst string
	fmt.Printf("\r富强|民主|文明|和谐|自由|平等|公正|法治|爱国|敬业|诚信|友善")
TRANS:
	result, fail := DeepLx.TranslateByDeepLX("auto", "zh", src, "")
	if fail != nil { //查询失败
		google, err := TransByGoogle(src, c, p)
		if err != nil {
			goto TRANS
		} else {
			c.SetGoogle()
			dst = google
		}
	} else {
		c.SetDeeplx()
		dst = result
	}
	if dst == "" {
		Translate(src, p, c)
	}
	return dst
}

func Trans(fp string, p *constant.Param, c *constant.Count) {
	// todo 翻译字幕
	r := seed.Intn(2000)
	//中间文件名
	srt := strings.Replace(fp, p.GetPattern(), "srt", 1)
	tmpname := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), strconv.Itoa(r), ".srt"}, "")
	before := util.ReadInSlice(srt)
	after, _ := os.OpenFile(tmpname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	for i := 0; i < len(before); i += 4 {
		if i+3 > len(before) {
			continue
		}
		after.WriteString(fmt.Sprintf("%s", before[i]))
		after.WriteString(fmt.Sprintf("%s", before[i+1]))
		src := before[i+2]
		afterSrc := replace.GetSensitive(src)
		var dst string
		if get, err := sql.GetDatabase().Hash().Get("translations", src); err == nil {
			dst = get.String()
			fmt.Println("find in cache")
			c.SetCache()
		} else {
			dst = Translate(afterSrc, p, c)
		}
		dst = replace.GetSensitive(dst)
		sql.GetDatabase().Hash().Set("translations", src, dst)
		log.Printf("\r文件名:%v", tmpname)
		log.Printf("\r原文:%v", src)
		log.Printf("\r译文:%v", dst)
		after.WriteString(fmt.Sprintf("%s", src))
		after.WriteString(fmt.Sprintf("%s", dst))
		after.WriteString(fmt.Sprintf("%s", before[i+3]))
		after.Sync()
	}
	origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
	exec.Command("cp", srt, origin).CombinedOutput()
	os.Rename(tmpname, srt)
}
