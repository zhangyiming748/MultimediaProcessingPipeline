package ytdlp

import (
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"log"
	"os/exec"
)

func DownloadVideo(uri, proxy,location string) (fp string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	name_cmd := exec.Command("yt-dlp", "--proxy", proxy, "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]/mp4", "--no-playlist", "--paths", location, "--get-filename", uri)
	name := util.GetVideoName(name_cmd)
	log.Printf("当前下载的文件标题:%s", name)
	download_cmd := exec.Command("yt-dlp", "--proxy",proxy, "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]/mp4", "--no-playlist", "--paths", location, uri)
	util.ExecCommand4Ytdlp(download_cmd)
	log.Printf("当前下载成功的文件标题:%s", name)
	return name
}

