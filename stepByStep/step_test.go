package stepbystep

import (
	l "Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/model"
	"Multimedia_Processing_Pipeline/sql"
	"Multimedia_Processing_Pipeline/translate"
	"log"
	"os"
	"strings"
	"testing"
)

func init() {
	l.SetLog()
	user := "root"
	password := "163453"
	host := "192.168.2.10"
	port := "3306"
	dbname := "Translate"
	sql.SetMysql(user, password, host, port, dbname)
	e1 := sql.GetMysql().Sync2(model.TranslateHistory{})
	if e1 != nil {
		log.Fatalf("数据库表更新失败%v有可能是数据库未正常连接\n", e1)
	} else {
		log.Printf("%+v表结构同步成功\n", model.TranslateHistory{})
	}
	e2 := sql.GetMysql().Sync2(model.YtdlpHistory{})
	if e2 != nil {
		log.Fatalf("数据库表更新失败%v有可能是数据库未正常连接\n", e2)
	} else {
		log.Printf("%+v表结构同步成功\n", model.YtdlpHistory{})
	}
}

// go test -v -timeout 10h -run TestDownloadAll
func TestDownloadAll(t *testing.T) {
	file := "/app/stepByStep/links.txt"
	proxy := "192.168.2.10:8889"
	location := "/app/stepByStep"
	links := ReadLinkToSlice(file)
	for _, link := range links {
		y := new(model.YtdlpHistory)
		y.Url = link
		if found, err := y.FindByUrl(); found {
			log.Printf("文件已经被下载过了,来源:%v\n", y.Host)
		} else if err != nil {
			log.Printf("查询数据库的时候发生错误:%v\n", err)
		} else {
			log.Printf("开始下载:%v\n", link)
			name := RunYtdlp(link, proxy, location)
			y.Title = name
			if host, e := os.Hostname(); e != nil {
				log.Printf("找不到主机名:%v\n", e)
				y.Host = "unknown"
			} else {
				y.Host = host
			}
			if strings.Contains(link, "pornhub") {
				y.Source = "pornhub"
				y.Key = strings.Split(link, "=")[1]
			}
			if strings.Contains(link, "xvideo") {
				y.Source = "xvideo"
				text := strings.Split(link, "/")
				y.Key = text[len(text)-1]
			}
			one, e := y.InsertOne()
			if e != nil {
				log.Printf("新条目写入数据库发生错误:%v\n", err)
			} else {
				log.Printf("%v新条目写入数据库\n", one)
			}
		}
	}
}

// go test -v -timeout 10h -run TestWhisperAll
func TestWhisperAll(t *testing.T) {
	root := "/app/stepByStep"
	model_name := "medium.en"
	model_path := "/app"
	video_language := "English"
	video_directory := "/app"
	videos, err := FindVideoFiles(root)
	if err != nil {
		log.Fatalln(err)
	}
	for _, video := range videos {
		GetSubtitle(video, model_name, model_path, video_language, video_directory)
	}
}

// go test -v -timeout 10h -run TestTranslateAll
func TestTranslateAll(t *testing.T) {
	root := "/app/stepByStep"
	files, err := FindSubtitleFiles(root)
	if err != nil {
		log.Fatalln(err)
	}
	for _, srt := range files {
		translate.Trans(srt)
	}
}
