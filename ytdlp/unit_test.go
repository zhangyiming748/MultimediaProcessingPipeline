package ytdlp

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func init() {

}

// go test -timeout 2000h -v -run TestYTdlp
func TestYTdlp(t *testing.T) {
	p := &constant.Param{
		VideosLocation: "/Users/zen/github/MultimediaProcessingPipeline/ytdlp",
		Language:       "English",
		Pattern:        "mp4",
		Model:          "medium.en",
		ToolsLocation:  "/Users/zen/github/MultimediaProcessingPipeline",
		Proxy:          "127.0.0.1:8889",
	}
	file, err := os.OpenFile("fail.list", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	log.SetLog(p)
	link := filepath.Join(p.GetVideosLocation(), "link.list")
	uris := util.ReadByLine(link)
	for _, uri := range uris {
		if link := DownloadVideo(uri, p.Proxy, p.VideosLocation); link == "" {
			file.WriteString(fmt.Sprintln(uri))
		}
	}
	file.Sync()
}
