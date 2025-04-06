package translateShell

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// go test -timeout 2000m -v -run TestTransAll
// docker run -dit --name trans -v /c/Users/zen/Github/MultimediaProcessingPipeline:/app -v /c/Users/zen/Videos/export/sdde:/data zhangyiming748/stand:latest bash
func TestGithub(t *testing.T) {
	src := "hello world"
	dst := TransByGithubDeepLX(src)
	fmt.Println(dst)
}
func getFiles(currentDir string) (filePaths []string) {
	err := filepath.Walk(currentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 检查是否是文件
		if !info.IsDir() {
			filePaths = append(filePaths, path) // 将文件的绝对路径添加到切片
		}
		return nil
	})

	if err != nil {
		fmt.Println("遍历目录失败:", err)
		return
	}

	// 打印所有文件的绝对路径
	for _, filePath := range filePaths {
		fmt.Println(filePath)
	}
	return filePaths
}
