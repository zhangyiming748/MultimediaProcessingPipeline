package translateShell

import (
	"Multimedia_Processing_Pipeline/replace"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

/*
接受绝对路径
*/
func TransName(filename string) string {
	base := filepath.Base(filename) // 从绝对路径中提取文件名部分
	dir := filepath.Dir(filename)
	ext := filepath.Ext(filename)             // 从文件名提取扩展名部分
	name := strings.Replace(base, ext, "", 1) //纯文件名
	//log.Printf("base name is %v\ndir is %v\next is %v\nname is %v\n", base, dir, ext, name)
	name = replaceEnglishSquareBrackets(name)
	name = replace.ChinesePunctuation(name)
	zh_cn, err := TransOnce(name, "192.168.1.20:8889")
	if err != nil {
		log.Println("translate err:", err)
		return filename
	}
	zh_cn = strings.Replace(zh_cn, "\n", "", -1)
	//fmt.Println("zh_ch is ", zh_cn)
	n_name := strings.Join([]string{zh_cn, ext}, "")
	n_path := strings.Join([]string{dir, n_name}, string(os.PathSeparator))
	//strings.Join([]string{zh_cn, ext}, "")
	return n_path
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
