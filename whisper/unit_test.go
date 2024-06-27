package whisper

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"testing"
)

func init() {

}

// go test -v -run TestWhisper
func TestWhisper(t *testing.T) {
	p := &constant.Param{
		Root:     "/home/zen/git/MultimediaProcessingPipeline/ytdlp",
		Language: "Russian",
		Pattern:  "mp4",
		Model:    "base",
		Location: "/home/zen/git/MultimediaProcessingPipeline/ytdlp",
		Proxy:    "192.168.1.20:8889",
	}
	log.SetLog(p)
	GetSubtitle("/home/zen/git/MultimediaProcessingPipeline/ytdlp/Дрочу на порно-историю из интернета [645dc2da8e6a3].mp4", p)
}
