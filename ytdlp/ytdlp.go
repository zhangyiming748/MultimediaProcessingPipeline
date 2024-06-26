package ytdlp

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"log"
	"os/exec"
)

func DownloadVideo(uri string, p constant.Param) (fp string, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	cmd := exec.Command("yt-dlp", "--proxy", p.GetProxy(), "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]", "--no-playlist", uri)
	msg := fmt.Sprintf("正在运行命令:%s", cmd.String())
	destination, err := util.ExecCommand4YtdlpDestination(cmd, msg)
	if err != nil {
		log.Fatalf("命令运行产生错误:%v\n", err)
	} else if destination == "" {
		log.Fatalf("视频下载后找不到标题信息,命令原文:%s\n", cmd.String())
	} else {
		log.Printf("当前下载成功的文件标题:%s", destination)
	}
	destination, err = replace.Rename(destination)
	return destination, nil
}
