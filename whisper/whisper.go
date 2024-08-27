package whisper

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

/*
生成字幕后返回字幕的绝对路径
*/
func GetSubtitle(fp string, p *constant.Param) string {
	cmd := exec.Command("whisper", fp, "--model", p.GetModel(), "--model_dir", p.GetLocation(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", p.GetLanguage(), "--output_dir", p.GetRoot(), "--verbose", "True")
	msg := fmt.Sprintf("正在处理的文件:%s", fp)
	err := util.ExecCommand(cmd, msg)
	if err != nil {
		log.Printf("当前字幕生成错误\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
	}
	fp = strings.Replace(fp, filepath.Ext(fp), ".srt", 1)
	//replace.RemoveTrailingNewlines(fp)
	return fp
}
