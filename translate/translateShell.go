package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/model"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/sql"
	"Multimedia_Processing_Pipeline/util"
	"errors"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zhangyiming748/MultiTranslatorUnifier/logic"
	_ "github.com/zhangyiming748/MultiTranslatorUnifier/logic"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
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
	m := logic.Trans(src, p.GetProxy(), "")
	for key, value := range m {
		switch key {
		case "Google":
			c.SetGoogle()
		case "Bing":
			c.SetBing()
		case "LinuxDo":
			c.SetDeeplx()
		case "Github":
			c.SetDeeplx()
		}
		dst = value
	}
	return dst
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
		afterSrc := replace.GetSensitive(src)
		var dst string
		if val, err := sql.GetLevelDB().Get([]byte(src), nil); err == nil {
			dst = string(val)
			fmt.Println("在缓存中找到")
			c.SetCache()
		} else {
			if errors.Is(err, leveldb.ErrNotFound) {
				fmt.Println("未在缓存中找到")
			}
			dst = Translate(afterSrc, p, c)
			randomNumber := util.GetSeed().Intn(401) + 100
			time.Sleep(time.Duration(randomNumber) * time.Millisecond) // 暂停 100 毫秒
		}
		dst = replace.GetSensitive(dst)
		if err := sql.GetLevelDB().Put([]byte(src), []byte(dst), nil); err != nil {
			fmt.Printf("缓存写入数据库错误:%v\n", err)
		}
		log.Printf("翻译之后序号\"%s\"时间\"%s\"正文\"%s\"空行\"%s\"\n", before[i], before[i+1], before[i+2], before[i+3])
		log.Printf("原文\"%s\"\t译文\"%s\"\n", src, dst)
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
