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
	// 设置默认时区为 Asia/Shanghai
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err) // 如果加载时区失败，则直接 panic
	}
	time.Local = loc
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
	if isCUDAAvailable() {
		cmd = exec.Command("whisper", fp, "--model", p.GetModel(), "--device", "cuda", "--model_dir", p.GetToolsLocation(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", p.GetLanguage(), "--output_dir", p.GetVideosLocation(), "--verbose", "True")
	} else {
		cmd = exec.Command("whisper", fp, "--model", p.GetModel(), "--model_dir", p.GetToolsLocation(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", p.GetLanguage(), "--output_dir", p.GetVideosLocation(), "--verbose", "True")
	}
	log.Printf("命令: %s\n", cmd.String())
	startTime := time.Now()
	msg := fmt.Sprintf("%v正在处理的文件:%s", time.Now().Format("2006-01-02 15:04:05"), fp)

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

func isCUDAAvailable() bool {
	// 使用 nvidia-smi 命令检查 CUDA 是否可用
	cmd := exec.Command("nvidia-smi")
	output, err := cmd.CombinedOutput()

	// 如果命令执行出错，或者输出中不包含 "NVIDIA-SMI"，则认为 CUDA 不可用
	if err != nil || !strings.Contains(string(output), "NVIDIA-SMI") {
		return false
	}

	return true
}
