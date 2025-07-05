package translateShell

import (
	"Multimedia_Processing_Pipeline/util"
	"encoding/json"
	"log"
	"os/exec"
	"strings"
)

func TransOnLocal(src, proxy string) (dst string) {
	//trans -brief -engine google -proxy 192.168.2.10:8889 :zh-CN 错误原文:exit status 1
	cmd := exec.Command("trans", "-brief", "-engine", "bing", ":zh-CN", src)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Bing翻译出现错误:%v,%v\n", string(output), err)
		cmd = exec.Command("trans", "-brief", "-engine", "google", "-proxy", proxy, ":zh-CN", src)
		output, err = cmd.CombinedOutput()
		if err != nil {
			log.Printf("翻译出现错误:%v,%v\n", string(output), err)
			return ""
		}
	}
	result := strings.Replace(string(output), "\n", "", 1)
	result = strings.Replace(result, "\r\n", "", 1)
	return result
}

func TransByServer(src string) (dst string) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	params := map[string]string{
		"text":        src,
		"source_lang": "auto",
		"target_lang": "zh",
	}
	host := "https://api.deeplx.org/DrkwqR4tE3DRyOseVibFah62BJXmcIryt4I9rTtzXTs/translate"
	log.Println(host)
	b, err := util.HttpPostJson(headers, params, host)
	if err != nil {
		return ""
	}
	log.Println(b)

	var d DeepLXTranslationResult
	if e := json.Unmarshal(b, &d); e != nil {
		return ""
	}
	log.Printf("%+v\n", d)
	return d.Data
}

const PREFIX = "https://api.deeplx.org"
const SUFFIX = "translate"

type DeepLXTranslationResult struct {
	Code         int      `json:"code"`
	ID           int64    `json:"id"`
	Message      string   `json:"message,omitempty"`
	Data         string   `json:"data"`         // The primary translated text
	Alternatives []string `json:"alternatives"` // Other possible translations
	SourceLang   string   `json:"source_lang"`
	TargetLang   string   `json:"target_lang"`
	Method       string   `json:"method"`
}
