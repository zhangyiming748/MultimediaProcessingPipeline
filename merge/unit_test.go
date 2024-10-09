package merge

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func init() {

}

// go test -v -run TestMerge
func TestMerge(t *testing.T) {
	root := "C:\\Users\\zen\\Videos\\export\\en"
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
