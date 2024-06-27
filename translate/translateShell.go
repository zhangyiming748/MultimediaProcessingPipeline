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
	fmt.Println("富强|民主|文明|和谐")
	fmt.Println("自由|平等|公正|法治")
	fmt.Println("爱国|敬业|诚信|友善")
TRANS:
	result, fail := DeepLx.TranslateByDeepLX("auto", "zh", src, "")
	if fail != nil {
		time.Sleep(time.Duration(seed.Intn(5)+1) * time.Second)
		log.Println("等待重试")
		Translate(src, p, c)
	} else if result == "" {
		time.Sleep(time.Duration(seed.Intn(5)+1) * time.Second)
		log.Println("等待重试")
		google, err := TransByGoogle(src, c, p)
		if err != nil {
			goto TRANS
		} else {
			dst = google
		}
	} else {
		dst = result
		c.SetDeeplx()
	}
	return dst
}

func Trans(fp string, p *constant.Param, c *constant.Count) {
	// todo 翻译字幕
	r := seed.Intn(2000)
	//中间文件名
	srt := strings.Replace(fp, p.GetPattern(), "srt", 1)
	tmpname := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), strconv.Itoa(r), ".srt"}, "")
	before := util.ReadByLine(srt)
	after, _ := os.OpenFile(tmpname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	for i := 0; i < len(before); i += 4 {
		if i+3 > len(before) {
			continue
		}
		after.WriteString(fmt.Sprintf("%s\n", before[i]))
		after.WriteString(fmt.Sprintf("%s\n", before[i+1]))
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
		log.Printf("文件名:%v\n原文:%v\n译文:%v\n", tmpname, src, dst)
		after.WriteString(fmt.Sprintf("%s\n", src))
		after.WriteString(fmt.Sprintf("%s\n", dst))
		after.WriteString(fmt.Sprintf("%s\n", before[i+3]))
		after.Sync()
	}
	origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
	exec.Command("cp", srt, origin).CombinedOutput()
	os.Rename(tmpname, srt)
}
