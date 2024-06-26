package ytdlp

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/util"
	"testing"
)

func init() {

}

// go test -v -run TestYTdlp
func TestYTdlp(t *testing.T) {
	p := &constant.Param{
		Root:     "/mnt/c/Users/zen/Github/Multimedia_Processing_Pipeline/ytdlp",
		Language: "English",
		Pattern:  "mp4",
		Model:    "base",
		Location: "/mnt/c/Users/zen/Github/Multimedia_Processing_Pipeline/ytdlp",
		Proxy:    "192.168.1.5:8889",
	}
	log.SetLog(p)
	uris := util.ReadByLine("/mnt/c/Users/zen/Github/Multimedia_Processing_Pipeline/ytdlp/test.list")
	for _, uri := range uris {
		DownloadVideo(uri, *p)
	}
}
