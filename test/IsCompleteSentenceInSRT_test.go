package t

import (
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func IsCompleteSentenceInSRT(srt string) {
	report, _ := os.OpenFile("report.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	subtitle := util.ReadInSlice(srt)
	for i := 0; i < len(subtitle); i += 4 {

		if i+3 > len(subtitle) {
			continue
		}
		no := subtitle[i]
		src := subtitle[i+2]
		no = strings.ReplaceAll(no, "\n", "")
		src = strings.ReplaceAll(src, "\n", "")
		if strings.HasSuffix(src, ",") {
			continue
		}
		if strings.HasSuffix(src, ".") {
			continue
		}
		if strings.HasSuffix(src, "!") {
			continue
		}
		if strings.HasSuffix(src, "?") {
			continue
		}
		if strings.HasPrefix(src, ",") {
			continue
		}
		if strings.HasPrefix(src, ".") {
			continue
		}
		if strings.HasPrefix(src, "!") {
			continue
		}
		if strings.HasPrefix(src, "?") {
			continue
		}
		if strings.HasSuffix(src, "，") {
			continue
		}
		if strings.HasSuffix(src, "。") {
			continue
		}
		if strings.HasSuffix(src, "！") {
			continue
		}
		if strings.HasSuffix(src, "？") {
			continue
		}
		if strings.HasPrefix(src, "，") {
			continue
		}
		if strings.HasPrefix(src, "。") {
			continue
		}
		if strings.HasPrefix(src, "！") {
			continue
		}
		if strings.HasPrefix(src, "？") {
			continue
		}
		report.WriteString(fmt.Sprintf("文件:%s的第%s句:%v可能不是完整的一句话\n", srt, no, src))
	}
}
func GetSRTFiles(dir string) ([]string, error) {
	var srtFiles []string

	// 使用 Walk 函数遍历目录
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查文件是否是 SRT 文件
		if !info.IsDir() && filepath.Ext(path) == ".srt" {
			srtFiles = append(srtFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return srtFiles, nil
}

// go test -timeout 2000m -v -run TestGetReport
func TestGetReport(t *testing.T) {
	files, _ := GetSRTFiles("/mnt/f/joi")
	t.Log(files)
	for _, file := range files {
		IsCompleteSentenceInSRT(file)
	}
}
