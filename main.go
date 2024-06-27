package main

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/sql"
	"Multimedia_Processing_Pipeline/util"
	"os"
)

func initConfig(p *constant.Param) {
	log.SetLog(p)
	sql.SetDatabase(p)
	util.ExitAfterRun()
}
func main() {
	p := &constant.Param{
		Root:     "/mnt/c/Users/zen/Github/FastYt-dlp/joi2",
		Language: "English",
		Pattern:  "mp4",
		Model:    "base",
		Location: "/mnt/c/Users/zen/Github/FastYt-dlp/joi2",
		Proxy:    "127.0.0.1:8889",
		Merge:    false,
	}
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
}
