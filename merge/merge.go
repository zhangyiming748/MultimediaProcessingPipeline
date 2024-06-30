package merge

import (
	"Multimedia_Processing_Pipeline/constant"
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

func MkvWithAss(file string, p *constant.Param) {
	srt := strings.Replace(file, p.Pattern, "srt", 1)
	if isExist(srt) {
		output := strings.Replace(file, p.GetPattern(), "_with_subtitle.mkv", 1)
		output = strings.Replace(output, ".", "", 1)
		par := FastMediaInfo.GetStandMediaInfo(file)
		width, _ := strconv.Atoi(par.Video.Width)
		height, _ := strconv.Atoi(par.Video.Height)
		log.Printf("获取到的分辨率:%vx%v\t", width, height)
		crf := FastMediaInfo.GetCRF("vp9", width, height)
		if crf == "" {
			crf = "31"
		}
		//cmd := exec.Command("ffmpeg", "-i", file, "-itsoffset", "1", "-i", srt, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-c:s", "ass", output)
		cmd := exec.Command("ffmpeg", "-i", file, "-i", srt, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-c:s", "ass", output)
		fmt.Printf("生成的命令: %s\n", cmd.String())
		msg := fmt.Sprintf("正在合成的视频:%s帧数%s", file, par.Video.FrameCount)
		if p.GetMerge() {
			err := util.ExecCommand(cmd, msg)
			if err != nil {
				log.Fatalf("合成视频%v命令执行失败:%v 退出", output, err)
			} else {
				os.Remove(file)
			}
		} else {
			fp := strings.Join([]string{p.GetRoot(), "merge.sh"}, string(os.PathSeparator))
			openFile, _ := os.OpenFile(fp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
			file = strings.Join([]string{"\"", file, "\""}, "")
			output = strings.Join([]string{"\"", output, "\""}, "")
			cmd = exec.Command("ffmpeg", "-i", file, "-i", srt, "-c:v", "libvpx-vp9", "-crf", crf, "-c:a", "libvorbis", "-ac", "1", "-c:s", "ass", output)
			openFile.WriteString(cmd.String())
			openFile.WriteString("\n")
			openFile.Sync()
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
