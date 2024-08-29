package ytdlp

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/util"
	"testing"
)

func init() {

}

// go test -timeout 24h -v -run TestYTdlp
func TestYTdlp(t *testing.T) {
	p := &constant.Param{
		Root:     "/data",
		Language: "English",
		Pattern:  "mp4",
		Model:    "base",
		Location: "/data",
		Proxy:    "192.168.1.20:8889",
	}
	log.SetLog(p)
	uris := util.ReadByLine("/app/ytdlp/pantyhose.list")
	for _, uri := range uris {
		DownloadVideo(uri, p)
	}
}
