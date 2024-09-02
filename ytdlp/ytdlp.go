package ytdlp

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func DownloadVideo(uri string, p *constant.Param) (fp string, err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	cmd := exec.Command("yt-dlp", "--proxy", p.GetProxy(), "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]/mp4", "--no-playlist", "--paths", p.GetRoot(), uri)
	msg := fmt.Sprintf("正在运行命令:%s", cmd.String())
	destination, err := util.ExecCommand4YtdlpDestination(cmd, msg)
	if err != nil {
		log.Fatalf("命令运行产生错误:%v\n", err)
	} else if destination == "" {
		log.Printf("视频下载后找不到标题信息,命令原文:%s\n", cmd.String())
		return "", nil
	} else {
		destination = strings.Replace(destination, filepath.Ext(destination), ".mp4", 1)
		log.Printf("当前下载成功的文件标题:%s", destination)
	}
	//destination = strings.Join([]string{p.GetRoot(), destination}, string(os.PathSeparator))
	destination = strings.ReplaceAll(destination, "\n", "")
	originName := destination
	destination = replaceEnglishSquareBrackets(destination)
	destination = replaceChineseRoundBrackets(destination)
	destination = replaceEnglishRoundBrackets(destination)
	destination = replaceChineseParentheses(destination)
	destination = removeSpaceBeforeExtension(destination)
	log.Printf("重命名前:%s\t后:%s\n", originName, destination)
	err = os.Rename(originName, destination)
	if err != nil {
		log.Printf("重命名失败")
		return originName, nil
	} else {
		log.Printf("重命名成功")
	}

	return destination, nil
}

/*
替换英文方括号
*/
func replaceEnglishSquareBrackets(input string) string {
	//input := "这是一个测试字符串[包含方括号内容]，请忽略这部分内容。"
	re := regexp.MustCompile(`\[[^\]]*?\]`)
	result := re.ReplaceAllString(input, "")
	return result
}

/*
替换中文方括号
*/
func replaceChineseRoundBrackets(input string) string {
	//input := "这是一个测试字符串【包含中文括号内容】，请忽略这部分内容。"
	re := regexp.MustCompile(`【[^】]*?】`)
	result := re.ReplaceAllString(input, "")
	return result
}

/*
替换英文圆括号
*/
func replaceEnglishRoundBrackets(input string) string {
	//input := "这是一个测试字符串(包含英文括号内容)，请忽略这部分内容。"
	re := regexp.MustCompile(`\([^\)]*?\)`)
	result := re.ReplaceAllString(input, "")
	return result
}

/*
替换中文圆括号
*/
func replaceChineseParentheses(input string) string {
	//input := "这是一个测试字符串(包含英文括号内容)，请忽略这部分内容。"
	re := regexp.MustCompile(`（[^\)]*?）`)
	result := re.ReplaceAllString(input, "")
	return result
}

/*
替换扩展名前面的空格
*/
func removeSpaceBeforeExtension(input string) string {
	output := strings.Replace(input, " .", ".", 1)
	return output
}
