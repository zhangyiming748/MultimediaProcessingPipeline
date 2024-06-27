package merge

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"testing"
)

func init() {

}

// go test -v -run TestMerge
func TestMerge(t *testing.T) {
	p := &constant.Param{
		Root:     "/home/zen/git/MultimediaProcessingPipeline/ytdlp",
		Language: "English",
		Pattern:  "mp4",
		Model:    "base",
		Location: "/home/zen/git/MultimediaProcessingPipeline/ytdlp",
		Proxy:    "192.168.1.20:8889",
	}
	log.SetLog(p)
	MkvWithAss("/home/zen/git/MultimediaProcessingPipeline/ytdlp/EDGE FOR ME -JOI [656dc0089a6eb].mp4", p)
}
