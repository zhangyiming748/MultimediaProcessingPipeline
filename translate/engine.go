package translateShell

import (
	"log"
	"strings"
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

