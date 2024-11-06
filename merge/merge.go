package merge

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/replace"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Mp4WithSrt(file string) {
	srt := strings.Replace(file, filepath.Ext(file), ".srt", 1)
	if isExist(srt) {
		output := strings.Replace(file, filepath.Ext(file), "_with_subtitle.mp4", 1)
		output = replace.ReplaceEnglishSquareBrackets(output)
		par := FastMediaInfo.GetStandMediaInfo(file)
		width, _ := strconv.Atoi(par.Video.Width)
		height, _ := strconv.Atoi(par.Video.Height)
		log.Printf("获取到的分辨率:%vx%v\t", width, height)
		cmd := exec.Command("ffmpeg", "-i", file, "-f", "srt", "-i", srt, "-c:v", "libx265", "-c:a", "libopus", "-map_chapters", "-1", "-ac", "1", "-c:s", "mov_text", output)
		fmt.Printf("生成的命令: %s\n", cmd.String())
		err := util.ExecCommandWithBar(cmd, par.Video.FrameCount)
		if err != nil {
			constant.Error(fmt.Sprintf("合成视频%v命令执行失败:%v 退出", output, err))
			return
		} else {
			os.Remove(file)
		}
	}
	countdown_with_exit(10)
}

func isExist(fp string) bool {
	_, err := os.Stat(fp)
	if os.IsNotExist(err) {
		log.Printf("%s 对应的字幕文件不存在\n", fp)
		return false
	} else {
		log.Printf("%s 对应的字幕文件存在\n", fp)
		return true
	}
}

func countdown_with_exit(t int) {
	for i := t; i > 0; i-- {
		fmt.Printf("\r上一个视频转换完成,等待%d秒", i)
		time.Sleep(1 * time.Second) // 暂停1秒
	}
	fmt.Println("\r上一个视频转换完成,等待时间结束!")
}
