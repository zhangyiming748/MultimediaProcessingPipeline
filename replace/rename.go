package replace

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// ToDo 重命名会莫名奇失败 同样操作 mv命令成功 但在代码里调mv会报错找不到文件
func Rename(src string) (dst string, err error) {
	dst = BracketsContent(src)
	//cleanName = strings.Replace(cleanName, "｜", "", -1)
	dst = ForFileName(dst)
	if err = os.Rename(src, dst); err != nil {
		var e = err.Error()
		msg := strings.Join([]string{e, fmt.Sprintf("src:%s重命名dst:%s\t失败:%v\n", src, dst, err)}, "")
		return src, errors.New(msg)
	} else {
		log.Printf("src:%s重命名dst:%s成功\n", src, dst)
	}
	return dst, nil
}
