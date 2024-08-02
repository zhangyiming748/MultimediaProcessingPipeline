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
		Root:     "/mnt/c/Users/zen/Github/Multimedia_Processing_Pipeline",
		Language: "English",
		Pattern:  "mp4",
		Model:    "medium",
		Location: "/mnt/c/Users/zen/Github/Multimedia_Processing_Pipeline",
		Proxy:    "192.168.1.20:8889",
	}
	log.SetLog(p)
	fps := []string{
		"/mnt/c/Users/zen/Github/Multimedia_Processing_Pipeline/Tom Cruise Terrifies James in 'Top Gun' Fighter Jet! [v1iZtBM23bY].mp4",
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
		"F:\\large",
	}
	for _, fp := range fps {
		GetSubtitle(fp, p)
		translateShell.Trans(fp, p, &constant.Count{})
	}
}
