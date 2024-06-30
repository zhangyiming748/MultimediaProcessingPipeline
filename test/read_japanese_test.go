package t

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadJapanese(t *testing.T) {
	file, err := os.OpenFile("/Users/zen/Github/MultimediaProcessingPipeline/test/NieR：Automata Fan Festival 12022 koncert [wX_SAi_ZcFQ].srt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}

func TestReadAllJapanese(t *testing.T) {
	decoder := japanese.ShiftJIS.NewDecoder()
	content, err := ioutil.ReadFile("utf8.srt")
	if err != nil {
		// 处理错误
		return
	}

	decodedContent, err := decoder.Bytes(content)
	if err != nil {
		// 处理错误
		return
	}
	fmt.Println(string(decodedContent))

}
func TestReadStand(t *testing.T) {
	file, err := os.ReadFile("utf8.srt")
	if err != nil {
		return
	}
	fmt.Println(string(file))
}
func TestReadInSlice(t *testing.T) {
	// 读取文件内容到字节切片中
	fileBytes, err := os.ReadFile("utf8.srt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 创建一个bufio.Reader对象
	reader := bufio.NewReader(bytes.NewReader(fileBytes))

	// 按行读取文件内容并存储到字符串切片中
	var lines []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lines = append(lines, line)
	}

	// 打印结果
	for i, line := range lines {
		fmt.Printf("第%d行: %s\n", i+1, line)
	}
}
