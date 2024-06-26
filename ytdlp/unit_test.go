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
		Root:     "/home/zen/git/MultimediaProcessingPipeline/ytdlp",
		Language: "English",
		Pattern:  "mp4",
		Model:    "base",
		Location: "/home/zen/git/MultimediaProcessingPipeline/ytdlp",
		Proxy:    "192.168.1.20:8889",
	}
	log.SetLog(p)
	uris := util.ReadByLine("/home/zen/git/MultimediaProcessingPipeline/ytdlp/test.list")
	for _, uri := range uris {
		DownloadVideo(uri, *p)
	}
}
