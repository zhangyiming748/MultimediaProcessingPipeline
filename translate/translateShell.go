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
	"sync"
	"time"
)

var (
	seed = rand.New(rand.NewSource(time.Now().Unix()))
)

const (
	TIMEOUT = 8 //second
	RETRY   = 3
)

type Result struct {
	From string // 来源
	Dst  string // 翻译内容
}

func Translate(src string, p *constant.Param, c *constant.Count) (dst string) {
	//trans -brief ja:zh "私の手の動きに合わせて|そう"
	ch := make(chan Result)
	var once sync.Once
	proxy := p.GetProxy()
	language := ":zh-CN"
	retry := 0
	if p.GetProxy() == "" {
		fmt.Println("富强|民主|文明|和谐")
		fmt.Println("自由|平等|公正|法治")
		fmt.Println("爱国|敬业|诚信|友善")
		dst, _ = DeepLx.TranslateByDeepLX("auto", "zh", src, "")
	} else {
		for {
			go TransByGoogle(proxy, language, src, ch, c, &once)
			go TransByBing(proxy, language, src, ch, c, &once)
			//使用同一个通道 传递结构体 标明来源
			var result Result
			select {
			case result = <-ch:
				if result.From == "google" {
					c.SetGoogle()
				} else if result.From == "bing" {
					c.SetBing()
				}
				dst = result.Dst
			case <-time.After(TIMEOUT * time.Second):
				dst, _ = DeepLx.TranslateByDeepLX("auto", "zh", src, "")
				log.Printf("trans超时,使用本地deepXL翻译结果:%v\n", dst)
				c.SetDeeplx()
			}
			if dst != "" {
				break
			} else {
				retry++
				log.Printf("查询结果为空retry:%v\n", retry)
			}
			if retry >= RETRY {
				log.Fatalln("达到重试次数后依旧失败,需要检查网络")
			}
		}
	}
	dst = replace.ChinesePunctuation(dst)
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
			var count int
			for !replace.Success(dst) {
				if count > RETRY {
					log.Fatalf("达到重试次数后依旧失败,需要检查网络,srt=%v\tdst=%v\n", afterSrc, dst)
				}
				log.Printf("查询失败\t重试%v\n", count)
				time.Sleep(1 * time.Second)
				dst = Translate(afterSrc, p, c)
				count++
			}
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
