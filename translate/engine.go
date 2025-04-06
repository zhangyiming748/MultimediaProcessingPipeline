package translateShell

import (
	"log"
	"strings"

	translate "github.com/OwO-Network/DeepLX/translate"
)

func TransByLinuxdoDeepLX(src, apikey string) (dst string) {
	//apikey := os.Getenv("LINUXDO")
	result, err := Req(src, apikey)
	result = strings.Replace(result, "\\r\\n", "", 1)
	result = strings.Replace(result, "\n", "", 1)
	result = strings.Replace(result, "\r\n", "", 1)
	if result == "" {
		return
	}
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("linuxdo 版本 deeplx 返回:%+v\n", result)
	return result
}

func TransByGithubDeepLX(src string) (dst string) {
	ret, err := translate.TranslateByDeepLX("auto", "zh", src, "", "", "")
	if err != nil {
		log.Fatalf("Github版本deeplx翻译报错%v\n", err)
	}
	log.Printf("GitHub 版本 deeplx 返回:%+v\n", ret)
	dst = ret.Data
	dst = strings.Replace(dst, "\\r\\n", "", 1)
	dst = strings.Replace(dst, "\n", "", 1)
	dst = strings.Replace(dst, "\r\n", "", 1)
	return dst
}
