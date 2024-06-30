package main

import (
	"Multimedia_Processing_Pipeline/constant"
	mylog "Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/sql"
	translateShell "Multimedia_Processing_Pipeline/translate"
	"Multimedia_Processing_Pipeline/util"
	"Multimedia_Processing_Pipeline/whisper"
	"Multimedia_Processing_Pipeline/ytdlp"
	"log"
	"os"
)

func initConfig(p *constant.Param) {
	if !util.IsExistCmd("ffmpeg") {
		log.Fatalln("ffmpeg未安装")
	}
	if !util.IsExistCmd("whisper") {
		log.Fatalln("whisper未安装")
	}
	if !util.IsExistCmd("trans") {
		log.Fatalln("trans未安装")
	}
	if !util.IsExistCmd("yt-dlp") {
		log.Fatalln("yt-dlp")
	}
	mylog.SetLog(p)
	sql.SetDatabase(p)
	util.ExitAfterRun()
	replace.SetSensitive(p)
}
func main() {
	p := new(constant.Param)
	p.Root = "/home/zen/git/MultimediaProcessingPipeline/ytdlp"
	p.Language = "japanese"
	p.Pattern = "webm"
	p.Model = "base"
	p.Location = "/home/zen/git/MultimediaProcessingPipeline/ytdlp"
	p.Proxy = "192.168.1.20:8889"
	p.Merge = false
	initConfig(p)
	if root := os.Getenv("root"); root != "" {
		p.SetRoot(root)
	}
	if language := os.Getenv("language"); language != "" {
		p.SetLanguage(language)
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
	if merge := os.Getenv("merge"); merge == "1" {
		p.Merge = true
	}
	video, err := ytdlp.DownloadVideo("https://youtu.be/wX_SAi_ZcFQ", p)
	if err != nil {
		return
	}
	whisper.GetSubtitle(video, p)
	c := new(constant.Count)
	translateShell.Trans(video, p, c)
}
