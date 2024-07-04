package whisper

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	translateShell "Multimedia_Processing_Pipeline/translate"
	"testing"
)

func init() {

}

// go test -v -run TestWhisper
func TestWhisper(t *testing.T) {
	p := &constant.Param{
		Root:     "/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD",
		Language: "Japanese",
		Pattern:  "mp3",
		Model:    "base",
		Location: "/Users/zen/Downloads",
		Proxy:    "192.168.1.20:8889",
	}
	log.SetLog(p)
	fps := []string{
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD1/01 - アーニャとベッキーのびっくり大作戦！.mp3",
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD1/02 - アーニャとベッキーのびっくり大作戦！ Cast Commentary.mp3",
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD2/01 - ブライア姉弟のスペシャルクッキング.mp3",
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD2/02 - ブライア姉弟のスペシャルクッキング Cast Commentary.mp3",
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD3/01 - ケーキを選んで世界平和⁉.mp3",
		"/Users/zen/Downloads/[ReinForce] Spy x Family (BDRip 1920x1080 x264 FLAC)/Extra/CD/CD3/02 - ケーキを選んで世界平和⁉ Cast Commentary.mp3",
	}
	for _, fp := range fps {
		GetSubtitle(fp, p)
	}
}

// go test -v -run TestWhisperAndTrans
func TestWhisperAndTrans(t *testing.T) {
	p := &constant.Param{
		Root:     "C:\\Users\\zen\\Downloads",
		Language: "Japanese",
		Pattern:  "mp4",
		Model:    "large-v3",
		Location: "C:\\Users\\zen\\Downloads",
		Proxy:    "192.168.1.20:8889",
	}
	log.SetLog(p)
	fps := []string{
		"C:\\Users\\zen\\Downloads\\sdde-712.mp4",
	}
	for _, fp := range fps {
		GetSubtitle(fp, p)
		translateShell.Trans(fp, p, &constant.Count{})
	}
}
