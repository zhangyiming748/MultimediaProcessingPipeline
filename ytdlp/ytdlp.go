package ytdlp

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/util"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func DownloadVideo(uri string, p *constant.Param) (fp string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	name_cmd := exec.Command("yt-dlp", "--proxy", p.GetProxy(), "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]/mp4", "--no-playlist", "--paths", p.GetRoot(), "--get-filename", uri)
	name := util.GetVideoName(name_cmd)
	log.Printf("当前下载的文件标题:%s", name)
	download_cmd := exec.Command("yt-dlp", "--proxy", p.GetProxy(), "-f", "bestvideo[height<=?1080]+bestaudio/best[height<=?1080]/mp4", "--no-playlist", "--paths", p.GetRoot(), uri)
	util.ExecCommand4Ytdlp(download_cmd)
	log.Printf("当前下载成功的文件标题:%s", name)
	return name
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
