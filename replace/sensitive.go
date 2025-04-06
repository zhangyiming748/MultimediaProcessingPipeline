package replace

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/model"
	"bufio"
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
	fp1 := strings.Join([]string{p.GetVideosLocation(), "sensitive.txt"}, string(os.PathSeparator))
	fp2 := strings.Join([]string{p.GetToolsLocation(), "sensitive.txt"}, string(os.PathSeparator))
	fp3 := "sensitive.txt"
	lines := []string{}
	if IsExist(fp1) {
		log.Printf("从视频目录%v中加载敏感词\n", fp1)
		lines = append(lines, fp1)
	} else if IsExist(fp2) {
		log.Printf("从程序目录%v中加载敏感词\n", fp2)
		lines = append(lines, fp2)
	} else if IsExist(fp3) {
		log.Printf("从程序目录%v中加载敏感词\n", fp3)
		lines = append(lines, fp3)
	} else {
		log.Println("没有找到敏感词文件")
	}
	for _, line := range lines {
		key := strings.Split(line, ":")[0]
		value := strings.Split(line, ":")[1]
		s := new(model.Sensitive)
		s.Before = key
		s.After = value
		if found, err := s.FindBySrc(); !found || err != nil {
			one, err := s.InsertOne()
			if err != nil {
				return
			} else {
				log.Printf("插入敏感词%v:%v成功:%v\n", key, value, one)
			}
		}
	}
	m := new(model.Sensitive).GetAll()
	for _, v := range m {
		Sensitive[v.Before] = v.After
	}
}

func readByLine(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		log.Println("按行读文件出错", err)
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
func IsExist(folderPath string) bool {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		log.Printf("文件夹:%v不存在\n", folderPath)
		return false
	} else {
		log.Printf("文件夹:%v存在\n", folderPath)
		return true
	}
}
