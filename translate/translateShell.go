package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/util"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	seed = rand.New(rand.NewSource(time.Now().Unix()))
)

const (
	TIMEOUT = 10 //second

)

func Translate(src string, p *constant.Param, c *constant.Count) (dst string) {
	return TransByServer(src)
}

func Trans(fp string, p *constant.Param, c *constant.Count) {
	// todo 翻译字幕
	r := seed.Intn(2000)
	//中间文件名
	//srt := strings.Replace(fp, p.GetPattern(), "srt", 1)
	//log.Printf("trans接受到的fp=%s\n", fp)
	//log.Printf("此时的p=%+v\n", p)

	srt := strings.Replace(fp, p.GetPattern(), "srt", 1)
	log.Printf("%v根据文件名:%s\t替换的字幕名:%s\n", p.GetPattern(), fp, srt)
	//log.Fatalf("%v根据文件名:%s\t替换的字幕名:%s\n", p.GetPattern(), fp, srt)
	tmpname := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), strconv.Itoa(r), ".srt"}, "")
	before := util.ReadInSlice(srt)
	fmt.Println(before)
	after, _ := os.OpenFile(tmpname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	defer func() {
		if err := recover(); err != nil {
			v := fmt.Sprintf("捕获到错误:%v\n", err)
			if strings.Contains(v, "index out of range") {
				fmt.Println("捕获到 index out of range 类型错误,忽略并继续执行重命名操作")
				{
					origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
					err1 := os.Rename(srt, origin)
					err2 := os.Rename(tmpname, srt)
					if err1 != nil || err2 != nil {
						constant.Warning(fmt.Sprintf("字幕文件重命名出现错误:%v\n", err))
					}
				}
				return
			} else {
				log.Fatalf("捕获到其他错误:%v\n", v)
			}
		}
	}()
	for i := 0; i < len(before); i += 4 {
		if i+3 > len(before) {
			continue
		}
		log.Printf("翻译之前序号\"%s\"时间\"%s\"正文\"%s\"空行\"%s\"\n", before[i], before[i+1], before[i+2], before[i+3])
		log.SetPrefix(before[i])
		after.WriteString(before[i])
		after.WriteString(before[i+1])
		src := before[i+2]
		src = strings.Replace(src, "\n", "", 1)
		src = strings.Replace(src, "\r\n", "", 1)
		afterSrc := replace.GetSensitive(src)
		var dst string

		dst = Translate(afterSrc, p, c)
		dst = strings.Replace(dst, "\n", "", -1)
		randomNumber := util.GetSeed().Intn(401) + 100
		time.Sleep(time.Duration(randomNumber) * time.Millisecond) // 暂停 100 毫秒
		dst = replace.GetSensitive(dst)

		fmt.Printf("翻译之后序号:\"%s\"时间:\"%s\"正文:\"%s\"空行:\"%s\"原文:\"%s\"\t译文\"%s\"\n", before[i], before[i+1], before[i+2], before[i+3], src, dst)
		after.WriteString(src)
		after.WriteString("\n")
		after.WriteString(dst)
		after.WriteString(before[i+3])
		after.WriteString(before[i+3])
		after.Sync()
	}
	after.Close()
	origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
	err1 := os.Rename(srt, origin)
	err2 := os.Rename(tmpname, srt)
	if err1 != nil || err2 != nil {
		constant.Warning(fmt.Sprintf("字幕文件重命名出现错误:%v:%v\n", err1, err2))
	}
}

func Req(src string) (string, error) {
	log.Printf("开始翻译:%s\n", src)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	params := map[string]string{
		"src": src,
	}
	host := "http://192.168.2.10:8192/api/v1/translate"

	b, err := util.HttpPostJson(headers, params, host)
	if err != nil {
		return "", err
	}
	log.Printf("%v\n", string(b))
	var r Res
	if UnmarshalErr := json.Unmarshal(b, &r); UnmarshalErr != nil {
		return "", UnmarshalErr
	}
	return r.Dst, err
}

type Res struct {
	Src  string `json:"src"`
	Dst  string `json:"dst"`
	From string `json:"from"`
}
