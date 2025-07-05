package stepbystep

import (
	l "Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/translate"
	"log"
	"testing"
)

func init() {
	l.SetLog()
}

// go test -v -timeout 10h -run TestDownloadAll
func TestDownloadAll(t *testing.T) {
	file := "/app/stepByStep/links.txt"
	proxy := "192.168.1.102:8889"
	location := "/app/stepByStep"
	links := ReadLinkToSlice(file)
	for _, link := range links {
		RunYtdlp(link, proxy, location)
	}
}

// go test -v -timeout 10h -run TestWhisperAll
func TestWhisperAll(t *testing.T) {
	root := "/app/stepByStep"
	model_name := "medium.en"
	model_path := "/app"
	video_language := "English"
	video_directory := "/app"
	videos, err := FindVideoFiles(root)
	if err != nil {
		log.Fatalln(err)
	}
	for _, video := range videos {
		GetSubtitle(video, model_name, model_path, video_language, video_directory)
	}
}

// go test -v -timeout 10h -run TestTranslateAll
func TestTranslateAll(t *testing.T) {
	root := "/app/stepByStep"
	files, err := FindSubtitleFiles(root)
	if err != nil {
		log.Fatalln(err)
	}
	for _, srt := range files {
		translate.Trans(srt)
	}
}
