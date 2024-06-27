package replace

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/sql"
	"Multimedia_Processing_Pipeline/util"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var Sensitive = map[string]string{}

func GetSensitive(str string) string {
	for k, v := range Sensitive {
		if strings.Contains(str, k) {
			str = strings.Replace(str, k, v, -1)
			log.Printf("替换生效\tbefore:%v\tafter:%v\t替换之后的完整句子:%v\n", k, v, str)
		}
	}
	return str
}

func SetSensitive(p *constant.Param) {
	fp1 := strings.Join([]string{p.GetRoot(), "sensitive.txt"}, string(os.PathSeparator))
	fp2 := "sensitive.txt"
	lines := []string{}
	if util.IsExist(fp1) {
		log.Printf("从视频目录%v中加载敏感词\n", fp1)
		lines = readByLine(fp1)
	}
	if util.IsExist(fp2) {
		log.Printf("从程序目录%v中加载敏感词\n", fp2)
		lines = readByLine(fp1)
	} else {
		log.Println("没有找到敏感词文件")
	}
	for _, line := range lines {
		before := strings.Split(line, ":")[0]
		after := strings.Split(line, ":")[1]
		log.Printf("写入敏感词:\tbefore:%v\tafter:%v\n", before, after)
		Sensitive[before] = after
		set, err := sql.GetDatabase().Hash().Set("sensitive", before, after)
		if err != nil {
			log.Printf("敏感词%v写入数据库失败%v\n", set, err)
		} else {
			log.Printf("敏感词%v写入数据库成功\n", set)
		}
	}
}

func readByLine(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		log.Println("按行读文件出错")
		return []string{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines, string(a))
	}
	return lines
}
