package replace

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func ChinesePunctuation(str string) string {
	str = strings.Replace(str, "。", ".", -1)
	str = strings.Replace(str, "，", ",", -1)
	str = strings.Replace(str, "《", "(", -1)
	str = strings.Replace(str, "》", ")", -1)
	str = strings.Replace(str, "【", "(", -1)
	str = strings.Replace(str, "】", ")", -1)
	str = strings.Replace(str, "（", "(", -1)
	str = strings.Replace(str, "）", ")", -1)
	str = strings.Replace(str, "「", "(", -1)
	str = strings.Replace(str, "」", ")", -1)
	str = strings.Replace(str, "+", "_", -1)
	str = strings.Replace(str, "`", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\u00A0", "", -1)
	str = strings.Replace(str, "\u0000", "", -1)
	str = strings.Replace(str, "·", "", -1)
	str = strings.Replace(str, "\uE000", "", -1)
	str = strings.Replace(str, "\u000D", "", -1)
	str = strings.Replace(str, "、", "", -1)
	//str = strings.Replace(str, "/", "", -1)
	str = strings.Replace(str, "！", "", -1)
	str = strings.Replace(str, "|", "", -1)
	str = strings.Replace(str, "｜", "", -1)
	str = strings.Replace(str, ":", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "&", "", -1)
	str = strings.Replace(str, "？", "", -1)
	str = strings.Replace(str, "(", "", -1)
	str = strings.Replace(str, ")", "", -1)
	str = strings.Replace(str, "-", "", -1)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "“", "", -1)
	str = strings.Replace(str, "”", "", -1)
	str = strings.Replace(str, "--", "", -1)
	str = strings.Replace(str, "_", "", -1)
	str = strings.Replace(str, "：", "", -1)
	str = strings.Replace(str, "\ufeff", "", -1)
	str = strings.Replace(str, "\n", "", 1)
	str = strings.Replace(str, "33 40  ", "", 1)
	str = strings.Replace(str, "33 40", "", 1)
	return str
}

/*
仅保留文件名中的 数字 字母 和 中文
*/
func ForFileName(name string) string {
	nStr := ""
	for _, v := range name {
		if Effective(string(v)) {
			// fmt.Printf("%d\t有效%v\n", i, string(v))
			nStr = strings.Join([]string{nStr, string(v)}, "")
		}
	}
	log.Printf("正则表达式匹配数字字母汉字:%v\n", nStr)
	return nStr
}
func Effective(s string) bool {
	if s == "-" {
		s = " "
	}
	if s == " " {
		return true
	}
	if s == "｜" {
		return false
	}
	if s == "：" {
		return false
	}
	if s == "." {
		return true
	}

	num := regexp.MustCompile(`\d`)          // 匹配任意一个数字
	letter := regexp.MustCompile(`[a-zA-Z]`) // 匹配任意一个字母
	char := regexp.MustCompile(`[\p{Han}]`)  // 匹配任意一个汉字
	if num.MatchString(s) || letter.MatchString(s) || char.MatchString(s) {
		return true
	}
	return false
}

/*
使用正则表达式替换以下类型的字符串中中括号内部任意字符为空的方法
*/
func BracketsContent(input string) string {
	re := regexp.MustCompile(` \[.*?\]`)
	output := re.ReplaceAllString(input, "")
	return output
}

/*
使用正则表达式替换以下类型的字符串中 “一个英文句号一个字母后面任意三位数字” 为空的方法
*/
func RealName(input string) string {
	re := regexp.MustCompile(`\.[a-zA-Z]\d{3}`)
	output := re.ReplaceAllString(input, "")
	fmt.Println(output)
	return output
}

/*
判断查询是否成功
*/
func Success(dst string) bool {
	if strings.Contains(dst, "\u001B") {
		return false
	}
	if strings.Contains(dst, "Showingtranslation") {
		return false
	}
	if strings.Contains(dst, "Connectiontimedout.RetryingIPv4connection") {
		return false
	}
	if strings.Contains(dst, "[WARNING]") {
		return false
	}
	if strings.Contains(dst, "Didyoumean") {
		return false
	}
	if strings.Contains(dst, "[22m") {
		return false
	}
	if strings.Contains(dst, "[33mm") {
		return false
	}
	return false
}
