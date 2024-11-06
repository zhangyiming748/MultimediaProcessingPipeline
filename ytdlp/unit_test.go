package ytdlp

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"os"
	"testing"
)

func init() {

}

// go test -timeout 2000h -v -run TestYTdlp
func TestYTdlp(t *testing.T) {
	p := &constant.Param{
		Root:     "/data",
		Language: "English",
		Pattern:  "mp4",
		Model:    "small",
		Location: "/data",
		Proxy:    "192.168.1.31:8889",
	}
	file, err := os.OpenFile("fail.list", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	log.SetLog(p)
	uris := util.ReadByLine("/data/ytdlp/test.list")
	for _, uri := range uris {
		if link := DownloadVideo(uri, p); link == "" {
			file.WriteString(fmt.Sprintln(uri))
		}
	}
	file.Sync()
}
