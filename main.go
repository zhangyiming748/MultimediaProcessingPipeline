package main

import (
	"Multimedia_Processing_Pipeline/constant"
	mylog "Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/merge"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/sql"
	translateShell "Multimedia_Processing_Pipeline/translate"
	"Multimedia_Processing_Pipeline/util"
	"Multimedia_Processing_Pipeline/whisper"
	"Multimedia_Processing_Pipeline/ytdlp"
	"github.com/fatih/color"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
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
	mylog.SetLog(p)
	sql.SetLevelDB(p)
	//util.ExitAfterRun()
	replace.SetSensitive(p)
}
func main() {
	p := new(constant.Param)
	p.Root = "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline"
	p.Language = "English"
	p.Pattern = "mp4"
	p.Model = "medium.en"
	p.Location = "C:\\Users\\zen\\Github\\MultimediaProcessingPipeline"
	p.Proxy = "192.168.1.8:8889"
	p.Merge = true
	p.Lines = strings.Join([]string{p.GetRoot(), "link.list"}, string(os.PathSeparator))
	initConfig(p)
	if root := os.Getenv("root"); root != "" {
		p.SetRoot(root)
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
		p.SetLocation(location)
	}
	if proxy := os.Getenv("proxy"); proxy != "" {
		p.SetProxy(proxy)
	}
	if mux := os.Getenv("merge"); mux == "1" {
		p.Merge = true
	}
	if lines := os.Getenv("lines"); lines != "" {
		p.SetLines(strings.Join([]string{p.GetRoot(), lines}, string(os.PathSeparator)))
	}
	c := new(constant.Count)
	lines := util.ReadByLine(p.GetLines())
	// 创建一个通道
	file_ch := make(chan string)
	wg := new(sync.WaitGroup)
	if p.GetMerge() {
		go merge.MergeByChannel(file_ch, wg)
	}
	for _, line := range lines {
		video := ytdlp.DownloadVideo(line, p)
		log.Printf("下载后的文件名为:%s\n", video)
		video = strings.Replace(video, "\n", "", 1)
		whisper.GetSubtitle(video, p)
		color.Red("开始翻译")
		translateShell.Trans(video, p, c)
		if p.GetMerge() {
			wg.Add(1)
			file_ch <- video
			//merge.Mp4WithSrt(video)
		}
	}
	wg.Wait()
}
