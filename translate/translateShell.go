package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/model"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/util"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const PREFIX = "https://api.deeplx.org"
const SUFFIX = "translate"

var (
	seed = rand.New(rand.NewSource(time.Now().Unix()))
)

const (
	TIMEOUT = 10 //second

)

func TransByLinuxdoDeepLX(src, apikey string) (dst string) {
	//apikey := os.Getenv("LINUXDO")
	result, err := Req(src, apikey)
	result = strings.Replace(result, "\\r\\n", "", 1)
	result = strings.Replace(result, "\n", "", 1)
	result = strings.Replace(result, "\r\n", "", 1)
	if result == "" {
		return
	}
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("linuxdo 版本 deeplx 返回:%+v\n", result)
	return result
}
func Translate(src string, p *constant.Param, c *constant.Count) (dst string) {
	return TransByLinuxdoDeepLX(src, p.LinuxDo)
}

func Trans(fp string, p *constant.Param, c *constant.Count) {
	var archives []model.TranslateHistory
	defer func() {
		all, err := new(model.TranslateHistory).InsertAll(archives)
		if err != nil {
			log.Printf("最终插入mysql失败:%v\n", err)
		}
		log.Printf("最终插入mysql成功:%v\n", all)
	}()

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
		afterSrc := replace.GetSensitive(src)
		var dst string
		behind := new(model.TranslateHistory)
		behind.Src = src
		if has, _ := behind.FindBySrc(); has {
			dst = behind.Dst
			fmt.Println("在缓存中找到")
			c.SetCache()
		} else {
			fmt.Println("未在缓存中找到")
			dst = Translate(afterSrc, p, c)
			dst = strings.Replace(dst, "\n", "", -1)
			randomNumber := util.GetSeed().Intn(401) + 100
			time.Sleep(time.Duration(randomNumber) * time.Millisecond) // 暂停 100 毫秒
		}
		dst = replace.GetSensitive(dst)
		if _, err := behind.InsertOne(); err != nil {
			fmt.Printf("缓存写入数据库错误:%v\n", err)
		}
		fmt.Printf("翻译之后序号:\"%s\"时间:\"%s\"正文:\"%s\"空行:\"%s\"原文\"%s\"\t译文\"%s\"\n", before[i], before[i+1], before[i+2], before[i+3], src, dst)
		after.WriteString(src)
		after.WriteString(dst)
		after.WriteString(before[i+3])
		after.WriteString(before[i+3])
		after.Sync()
		archive := model.TranslateHistory{
			Src: src,
			Dst: dst,
		}
		archives = append(archives, archive)
	}
	after.Close()
	origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
	err1 := os.Rename(srt, origin)
	err2 := os.Rename(tmpname, srt)
	if err1 != nil || err2 != nil {
		constant.Warning(fmt.Sprintf("字幕文件重命名出现错误:%v:%v\n", err1, err2))
	}
}

func TransFile(input string, p *constant.Param) {
	//translate-shell -i input.txt -o output.txt -t zh-CN
	output := strings.Replace(input, filepath.Ext(input), "_zhCN.txt", 1)
	cmd := exec.Command("translate-shell", "-e", "google", "-x", p.GetProxy(), "-i", input, "-o", output, "-t", "zh-CN")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("连接Stdout产生错误:%v\n", err)
		return
	}
	cmd.Stderr = cmd.Stdout
	if err = cmd.Start(); err != nil {
		log.Printf("启动cmd命令产生错误:%v\n", err)
		return
	}
	go func() {
		for {
			tmp := make([]byte, 1024)
			_, err := stdout.Read(tmp)
			t := string(tmp)
			t = strings.Replace(t, "\u0000", "", -1)
			fmt.Print(t)
			if err != nil {
				break
			}
		}
	}()
	if err = cmd.Wait(); err != nil {
		log.Printf("命令执行中产生错误:%v\n", err)
		return
	}
}
func Req(src, apikey string) (string, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	params := map[string]string{
		"text":        src,
		"source_lang": "auto",
		"target_lang": "zh",
	}
	host := strings.Join([]string{PREFIX, apikey, SUFFIX}, "/")

	b, err := util.HttpPostJson(headers, params, host)
	if err != nil {
		return "", err
	}
	log.Printf("%v\n", string(b))
	var d DeepLXTranslationResult
	if err := json.Unmarshal(b, &d); err != nil {
		return "", err
	}
	return d.Data, err
}

type DeepLXTranslationResult struct {
	Code         int      `json:"code"`
	ID           int64    `json:"id"`
	Message      string   `json:"message,omitempty"`
	Data         string   `json:"data"`         // The primary translated text
	Alternatives []string `json:"alternatives"` // Other possible translations
	SourceLang   string   `json:"source_lang"`
	TargetLang   string   `json:"target_lang"`
	Method       string   `json:"method"`
}
