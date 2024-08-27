package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/sql"
	"Multimedia_Processing_Pipeline/util"
	"errors"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
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
	fmt.Println("富强|民主|文明|和谐|自由|平等|公正|法治|爱国|敬业|诚信|友善")
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
		time.Sleep(1 * time.Second)
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
	//srt := strings.Replace(fp, p.GetPattern(), "srt", 1)
	log.Printf("trans接受到的fp=%s\n", fp)
	log.Printf("此时的p=%+v\n", p)

	srt := strings.Replace(fp, p.GetPattern(), "srt", 1)
	log.Printf("%v根据文件名:%s\t替换的字幕名:%s\n", p.GetPattern(), fp, srt)
	//log.Fatalf("%v根据文件名:%s\t替换的字幕名:%s\n", p.GetPattern(), fp, srt)
	tmpname := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), strconv.Itoa(r), ".srt"}, "")
	var before []string
	before = util.ReadInSlice(srt)
	after, _ := os.OpenFile(tmpname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	for i := 0; i < len(before); i += 4 {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err) // 处理错误
			}
		}()
		if i+3 > len(before) {
			continue
		}
		after.WriteString(fmt.Sprintf("%s", before[i]))
		after.WriteString(fmt.Sprintf("%s", before[i+1]))
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
		}
		dst = replace.GetSensitive(dst)
		if err := sql.GetLevelDB().Put([]byte(src), []byte(dst), nil); err != nil {
			log.Printf("缓存写入数据库错误:%v\n", err)
		}
		fmt.Printf("文件名:%v\t原文:%v\t译文:%v\n", tmpname, src, dst)
		after.WriteString(fmt.Sprintf("%s", src))
		after.WriteString(fmt.Sprintf("%s", dst))
		after.WriteString(fmt.Sprintf("%s", before[i+3]))
		after.Sync()
	}
	origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
	exec.Command("cp", srt, origin).CombinedOutput()
	os.Rename(tmpname, srt)
}
