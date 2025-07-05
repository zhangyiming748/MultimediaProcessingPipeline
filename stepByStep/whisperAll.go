package stepbystep

import (
	"github.com/h2non/filetype"
	"log"
	"os"
	"time"
	"path/filepath"
	"os/exec"
	"fmt"
	"strings"
	"Multimedia_Processing_Pipeline/util"
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

// FindFiles finds all files (not directories) in a given path, recursively.
// It's similar to the `find . -type f` command.
// It returns a slice of strings with the absolute paths of the files.
func FindVideoFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if it is a regular file.
		if !info.IsDir() {
			absPath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			if isVideo(absPath) {
				files = append(files, absPath)
			}

		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func isVideo(absPath string) bool {
	// Open a file descriptor
	file, _ := os.Open(absPath)
	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)
	if filetype.IsVideo(head) {
		log.Printf("%v is a video\n", absPath)
		return true
	} else {
		log.Printf("%v Not a video\n", absPath)
		return false
	}
}

/*
生成字幕后返回字幕的绝对路径
*/
func GetSubtitle(fp, model_name, model_path, video_language, video_directory string) string {
	err := os.Setenv("PYTHONIOENCODING", "utf-8")
	if err == nil {
		log.Println("utf-8环境设置成功")
	}
	var cmd *exec.Cmd
	if isCUDAAvailable() {
		cmd = exec.Command("whisper", fp, "--model", model_name, "--device", "cuda", "--model_dir", model_path, "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", video_language, "--output_dir", video_directory, "--verbose", "True")
	} else {
		cmd = exec.Command("whisper", fp, "--model", model_name, "--model_dir", model_path, "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", video_language, "--output_dir", video_directory, "--verbose", "True")
	}
	log.Printf("命令: %s\n", cmd.String())
	startTime := time.Now()
	msg := fmt.Sprintf("%v正在处理的文件:%s", time.Now().Format("2006-01-02 15:04:05"), fp)

	err = util.ExecCommand(cmd, msg)
	if err != nil {
		log.Printf("当前字幕生成错误\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
	}
	fp = strings.Replace(fp, filepath.Ext(fp), ".srt", 1)

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
