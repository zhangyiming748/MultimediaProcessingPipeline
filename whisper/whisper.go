package whisper

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	os.Setenv("PYTHONWARNINGS", "ignore::FutureWarning")
}

/*
生成字幕后返回字幕的绝对路径
*/
func GetSubtitle(fp string, p *constant.Param, fast bool) string {
	err := os.Setenv("PYTHONIOENCODING", "utf-8")
	if err == nil {
		log.Println("utf-8环境设置成功")
	}
	var cmd *exec.Cmd
	if hostname, unknown := os.Hostname(); unknown != nil {
		fmt.Println("未找到计算机名")
		cmd = exec.Command("whisper", fp, "--model", p.GetModel(), "--model_dir", p.GetToolsLocation(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", p.GetLanguage(), "--output_dir", p.GetVideosLocation(), "--verbose", "True")
	} else if hostname == constant.HASEE {
		fmt.Println("是神舟战神,可以使用cuda加速")
		cmd = exec.Command("whisper", fp, "--model", p.GetModel(), "--device", "cuda", "--model_dir", p.GetToolsLocation(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", p.GetLanguage(), "--output_dir", p.GetVideosLocation(), "--verbose", "True")
	} else {
		fmt.Println("是其他电脑,使用cpu硬肝")
		cmd = exec.Command("whisper", fp, "--model", p.GetModel(), "--model_dir", p.GetToolsLocation(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", p.GetLanguage(), "--output_dir", p.GetVideosLocation(), "--verbose", "True")
	}

	startTime := time.Now()
	msg := fmt.Sprintf("正在处理的文件:%s", fp)

	if fast {
		return cmd.String()
	} else {
		err = util.ExecCommand(cmd, msg)
		if err != nil {
			log.Printf("当前字幕生成错误\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		}
		fp = strings.Replace(fp, filepath.Ext(fp), ".srt", 1)
	}

	//replace.RemoveTrailingNewlines(fp)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	totalMinutes := duration.Seconds() / 60
	log.Printf("文件%v\n总共用时: %.2f 分钟\n", fp, totalMinutes)
	return fp
}
