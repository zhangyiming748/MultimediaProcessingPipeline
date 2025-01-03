package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/model"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/sql"
	"Multimedia_Processing_Pipeline/util"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	seed = rand.New(rand.NewSource(time.Now().Unix()))
)

const (
	TIMEOUT = 10 //second

)

func Translate(src string, p *constant.Param, c *constant.Count) string {
	//trans -brief ja:zh "私の手の動きに合わせて|そう"
	var dst string
	if src == "" {
		return dst
	}
	//fmt.Println("富强|民主|文明|和谐|自由|平等|公正|法治|爱国|敬业|诚信|友善")
	once := new(sync.Once)
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	ack := make(chan string, 1)
	wg.Add(1)
	//go TransByDeeplx(src, p.GetProxy(), once, wg, ack)
	if runtime.GOOS == "windows" {
		go TransByDeeplx(src, p, once, wg, ack)
	} else {
		go TransByGoogle(src, p.GetProxy(), once, wg, ack)
		go TransByBing(src, p.GetProxy(), once, wg, ack)
		go TransByDeeplx(src, p, once, wg, ack)
	}
	select {
	case dst = <-ack:
		//constant.Info(fmt.Sprintf("收到翻译结果:%v\n", dst))
	case <-time.After(TIMEOUT * time.Second): // 设置超时时间为5秒
		fmt.Printf("翻译超时,重试\n此时的src = %v\n", src)
		Translate(src, p, c)
	}
	if dst == "" {
		fmt.Printf("翻译结果为空,重试\n此时的src = %v\n", src)
		return src
	}
	dst = strings.Replace(dst, "\r\n", "", -1)
	dst = strings.Replace(dst, "\\n", "", -1)
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
