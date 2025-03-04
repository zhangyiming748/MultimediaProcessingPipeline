package translateShell

import (
	"regexp"
)

/*
替换英文方括号
*/
func replaceEnglishSquareBrackets(input string) string {
	//input := "这是一个测试字符串[包含方括号内容]，请忽略这部分内容。"
	re := regexp.MustCompile(`\[[^\]]*?\]`)
	result := re.ReplaceAllString(input, "")
	return result
}
