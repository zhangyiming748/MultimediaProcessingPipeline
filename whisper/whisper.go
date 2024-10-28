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
func GetSubtitle(fp string, p *constant.Param) string {
	err := os.Setenv("PYTHONIOENCODING", "utf-8")
	if err == nil {
		log.Println("utf-8环境设置成功")
	}

	cmd := exec.Command("whisper", fp, "--model", p.GetModel(), "--model_dir", p.GetLocation(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", p.GetLanguage(), "--output_dir", p.GetRoot(), "--verbose", "True")
	//cmd := exec.Command("whisper",  fp, "--model", p.GetModel(), "--model_dir", p.GetLocation(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", p.GetLanguage(), "--output_dir", p.GetRoot(), "--verbose", "True")
	startTime := time.Now()
	log.Printf("文件%v开始时间%v", fp, startTime.Format("20060102 15:04:05"))
	msg := fmt.Sprintf("正在处理的文件:%s", fp)
	err = util.ExecCommand(cmd, msg)
	if err != nil {
		log.Printf("当前字幕生成错误\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
	}
	fp = strings.Replace(fp, filepath.Ext(fp), ".srt", 1)
	//replace.RemoveTrailingNewlines(fp)
	endTime := time.Now()
	log.Printf("文件%v开始时间%v", fp, endTime.Format("20060102 15:04:05"))
	duration := endTime.Sub(startTime)
	totalMinutes := duration.Seconds() / 60
	log.Printf("文件%v总共用时: %.2f 分钟\n", fp, totalMinutes)
	return fp
}
