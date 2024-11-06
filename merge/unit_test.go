package merge

import (
	"Multimedia_Processing_Pipeline/replace"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func init() {

}

// go test -v -run TestMerge
func TestMerge(t *testing.T) {
	root := "E:\\Downloads\\My Pack\\GVG-170\\Pack From Shared\\ReADA"
	mp4s, _ := getMP4Files(root)
	for _, mp4 := range mp4s {
		Mp4WithSrt(mp4)
	}
}

// getMP4Files 遍历指定目录，返回所有 mp4 文件的路径
func getMP4Files(dir string) ([]string, error) {
	var mp4Files []string

	// 使用 Walk 函数遍历目录
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查文件是否是 mp4 文件
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".mp4") {
			mp4Files = append(mp4Files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return mp4Files, nil
}

func TestReplaceEnglish(t *testing.T) {
	input := "这是一个测试字符串 [包含方括号内容],请忽略这部分内容。"
	ret := replace.ReplaceEnglishSquareBrackets(input)
	t.Log(ret)
}
