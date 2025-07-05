package stepbystep

import (
	"log"
	"testing"
)

// go test -v -timeout 10h -run TestDownloadAll
func TestDownloadAll(t *testing.T) {
	file := "/app/stepByStep/links.txt"
	proxy := "127.0.0.1:8889"
	location := "/app/stepByStep"
	links := ReadLinkToSlice(file)
	for _, link := range links {
		RunYtdlp(link, proxy, location)
	}
}

// go test -v -timeout 10h -run TestWhisperAll
func TestWhisperAll(t *testing.T) {
	root := "/app/stepByStep/links.txt"
	model_name := ""
	model_path := ""
	video_language := ""
	video_directory := ""
	videos, err := FindVideoFiles(root)
	if err != nil {
		log.Fatalln(err)
	}
	for _, video := range videos {
		GetSubtitle(video, model_name, model_path, video_language, video_directory)
	}
}
