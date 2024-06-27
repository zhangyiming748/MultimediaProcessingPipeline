package merge

import (
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func MkvWithAss(file string, pattern string) {
	srt := strings.Replace(file, pattern, ".srt", 1)
	if isExist(srt) {
		output := strings.Replace(file, pattern, "_with_subtitle.mkv", 1)
		p := FastMediaInfo.GetStandMediaInfo(file)
		width, _ := strconv.Atoi(p.Video.Width)
		height, _ := strconv.Atoi(p.Video.Height)
		log.Printf("获取到的分辨率:%vx%v\t", width, height)
		crf := FastMediaInfo.GetCRF("vp9", width, height)
		if crf == "" {
			crf = "31"
		}
		//cmd := exec.Command("ffmpeg", "-i", file, "-itsoffset", "1", "-i", srt, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-c:s", "ass", output)
		cmd := exec.Command("ffmpeg", "-i", file, "-i", srt, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-c:s", "ass", output)
		fmt.Printf("生成的命令: %s\n", cmd.String())
		msg := fmt.Sprintf("正在合成的视频:%s", file)
		err := util.ExecCommand(cmd, msg)
		if err != nil {
			log.Fatalf("合成视频命令执行失败:%v 退出:%v", err)
		} else {
			os.Remove(file)
		}
	}
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
func getFilesWithExtension(folderPath string, extension string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
