package replace

import (
	"bufio"
	"os"
)

/*
它会读取文件内容，删除末尾的换行符，并将结果写回原文件。
*/
func RemoveTrailingNewlines(filePath string) error {
	inputFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	var outputBuffer []byte
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		outputBuffer = append(outputBuffer, scanner.Bytes()...)
		outputBuffer = append(outputBuffer, '\n')
	}

	// 移除末尾的换行符
	for len(outputBuffer) > 0 && outputBuffer[len(outputBuffer)-1] == '\n' {
		outputBuffer = outputBuffer[:len(outputBuffer)-1]
	}

	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = outputFile.Write(outputBuffer)
	if err != nil {
		return err
	}

	return nil
}
