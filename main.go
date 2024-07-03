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
	"log"
	"os"
	"path"
	"strings"
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
	sql.SetLevelDB(p)
	//util.ExitAfterRun()
	replace.SetSensitive(p)
}
func main() {
	p := new(constant.Param)
	p.Root = "/data"
	p.Language = "English"
	p.Pattern = "mp4"
	p.Model = "base"
	p.Location = "/data"
	p.Proxy = "192.168.1.20:8889"
	p.Merge = false
	p.Lines = strings.Join([]string{p.GetRoot(), "link.list"}, string(os.PathSeparator))
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
	if mux := os.Getenv("merge"); mux == "1" {
		p.Merge = true
	}
	if lines := os.Getenv("lines"); lines != "" {
		p.SetLines(strings.Join([]string{p.GetRoot(), lines}, string(os.PathSeparator)))
	}
	c := new(constant.Count)
	lines := util.ReadByLine(p.GetLines())
	for _, line := range lines {
		video, err := ytdlp.DownloadVideo(line, p)
		if err != nil {
			return
		}
		ext := strings.Replace(path.Ext(video), ".", "", 1)
		p.SetPattern(ext)
		log.Printf("下载后的文件名为:%s\n", video)
		whisper.GetSubtitle(video, p)
		translateShell.Trans(video, p, c)
		merge.MkvWithAss(video, p)
	}
}
