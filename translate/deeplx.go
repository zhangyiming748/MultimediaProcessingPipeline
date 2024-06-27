package translateShell

import (
	"Multimedia_Processing_Pipeline/util"
	"encoding/json"
	"log"
)

/*
curl --location 'https://api.deeplx.org/translate' \
--header 'Content-Type: application/json' \

	--data '{
	    "text": "Hello, world!",
	    "source_lang": "EN",
	    "target_lang": "ZH"
	}'
*/
type ans struct {
	Alternatives []string `json:"alternatives"`
	Code         int      `json:"code"`
	Data         string   `json:"data"`
	Id           int64    `json:"id"`
	Method       string   `json:"method"`
	SourceLang   string   `json:"source_lang"`
	TargetLang   string   `json:"target_lang"`
}

func DeepXl(src string) string {
	uri := "http://192.168.1.6:1188/translate"
	data := map[string]string{
		"text":        src,
		"source_lang": "auto",
		"target_lang": "ZH",
	}
	b, _ := util.HttpPostJson(nil, data, uri)

	var a ans

	json.Unmarshal(b, &a)
	if a.Data == "" {
		log.Fatalln("翻译结果为空,是否开启了全局代理")
	}
	return a.Data
}
