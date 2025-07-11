package main

import (
	"Multimedia_Processing_Pipeline/constant"
	mylog "Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/sql"
	stepbystep "Multimedia_Processing_Pipeline/stepByStep"
	translateShell "Multimedia_Processing_Pipeline/translate"
	"Multimedia_Processing_Pipeline/util"
	"Multimedia_Processing_Pipeline/ytdlp"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/fatih/color"
)

func initConfig(p *constant.Param) {
	if !util.IsExistCmd("ffmpeg") {
		log.Fatalln("ffmpeg未安装")
	}
	if !util.IsExistCmd("whisper") {
		log.Fatalln("whisper未安装")
	}
	if !util.IsExistCmd("trans") {
		if runtime.GOOS != "windows" {
			log.Fatalln("trans未安装")
		}
	}
	if !util.IsExistCmd("yt-dlp") {
		log.Fatalln("yt-dlp")
	}
	mylog.SetLog()
	sql.SetMysql("root", "163453", "192.168.2.10", "3306", "Translate")
	//util.ExitAfterRun()
	replace.SetSensitive(p)
	log.SetFlags(log.Ltime | log.Lshortfile)
}
func main() {
	p := new(constant.Param)
	p.VideosLocation = "/data"
	p.Language = "English"
	p.Pattern = "mp4"
	p.Model = "medium.en"
	p.ToolsLocation = "/data"
	p.Proxy = "192.168.1.31:8889"
	p.Merge = false
	p.Lines = strings.Join([]string{p.GetVideosLocation(), "link.list"}, string(os.PathSeparator))
	initConfig(p)
	if root := os.Getenv("root"); root != "" {
		p.SetVideosLocation(root)
	}
	if language := os.Getenv("language"); language != "" {
		p.SetLanguage(language)
	}
	if language := os.Getenv("language"); language == "English" {
		p.SetLanguage("medium.en")
	}
	if pattern := os.Getenv("pattern"); pattern != "" {
		p.SetPattern(pattern)
	}
	if model := os.Getenv("model"); model != "" {
		p.SetModel(model)
	}
	if location := os.Getenv("location"); location != "" {
		p.SetToolsLocation(location)
	}
	if proxy := os.Getenv("proxy"); proxy != "" {
		p.SetProxy(proxy)
	}
	if mux := os.Getenv("merge"); mux == "1" {
		p.Merge = true
	}
	if lines := os.Getenv("lines"); lines != "" {
		p.SetLines(strings.Join([]string{p.GetVideosLocation(), lines}, string(os.PathSeparator)))
	}
	lines := util.ReadByLine(p.GetLines())
	// 创建一个通道
	file_ch := make(chan string)
	wg := new(sync.WaitGroup)
	for _, line := range lines {
		video := ytdlp.DownloadVideo(line, p.GetProxy(), p.GetVideosLocation())
		log.Printf("下载后的文件名为:%s\n", video)
		video = strings.Replace(video, "\n", "", 1)
		stepbystep.GetSubtitle(video, p.GetModel(), p.GetToolsLocation(), p.GetLanguage(), p.GetVideosLocation())
		color.Red("开始翻译")
		translateShell.Trans(video)
		if p.GetMerge() {
			wg.Add(1)
			file_ch <- video
			//merge.Mp4WithSrt(video)
		}
	}
	wg.Wait()
}
