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
	p := new(constant.Param)
	p.Root = "/home/zen/git/MultimediaProcessingPipeline/ytdlp"
	p.Language = "Japanese"
	p.Pattern = "webm"
	p.Model = "base"
	p.Location = "/home/zen/git/MultimediaProcessingPipeline/ytdlp"
	p.Proxy = "192.168.1.20:8889"
	p.Merge = false
	log.SetLog(p)
	GetSubtitle("/home/zen/git/MultimediaProcessingPipeline/ytdlp/NieR：Automata Fan Festival 12022 koncert [wX_SAi_ZcFQ].webm", p)
}

//whisper "/home/zen/git/MultimediaProcessingPipeline/ytdlp/NieR：Automata Fan Festival 12022 koncert [wX_SAi_ZcFQ].webm" --model base --model_dir /home/zen/git/MultimediaProcessingPipeline/ytdlp --output_format srt --prepend_punctuations ,.? --language japanese --output_dir /home/zen/git/MultimediaProcessingPipeline/ytdlp --verbose True
