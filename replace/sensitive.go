package replace

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/sql"
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
	fp1 := strings.Join([]string{p.GetRoot(), "sensitive.txt"}, string(os.PathSeparator))
	fp2 := strings.Join([]string{p.GetLocation(), "sensitive.txt"}, string(os.PathSeparator))
	fp3 := "sensitive.txt"
	lines := []string{}
	if IsExist(fp1) {
		log.Printf("从视频目录%v中加载敏感词\n", fp1)
		lines = readByLine(fp1)
	} else if IsExist(fp2) {
		log.Printf("从程序目录%v中加载敏感词\n", fp2)
		lines = readByLine(fp2)
	} else if IsExist(fp3) {
		log.Printf("从程序目录%v中加载敏感词\n", fp3)
		lines = readByLine(fp3)
	} else {
		log.Println("没有找到敏感词文件")
	}
	//batch := new(leveldb.Batch)
	//for _, line := range lines {
	//	key := strings.Split(line, ":")[0]
	//	value := strings.Split(line, ":")[1]
	//	batch.Put([]byte(key), []byte(value))
	//}
	//err := sql.GetLevelDB().Write(batch, nil)
	//if err !=nil{
	//	log.Fatalf("敏感词写入数据库失败:%v\n",err)
	//}
	for _, line := range lines {
		before := strings.Split(line, ":")[0]
		after := strings.Split(line, ":")[1]
		log.Printf("写入敏感词 : before - %v\tafter - %v\n", before, after)
		Sensitive[before] = after
		err := sql.GetLevelDB().Put([]byte(before), []byte(after), nil)
		if err != nil {
			log.Printf("敏感词:%v - %v写入数据库失败:%v\n", before, after, err)
		} else {
			log.Printf("敏感词:%v - %v写入数据库成功\n", before, after)
		}
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
