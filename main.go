package main

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/sql"
	"Multimedia_Processing_Pipeline/util"
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
	}
	initConfig(p)
}
