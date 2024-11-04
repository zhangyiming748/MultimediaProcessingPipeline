package translateShell

import (
	"Multimedia_Processing_Pipeline/util"
	"encoding/json"
	"errors"
	"fmt"
	DP "github.com/OwO-Network/DeepLX/translate"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type DeeplxRep struct {
	Alternatives []string `json:"alternatives"`
	Code         int      `json:"code"`
	Data         string   `json:"data"`
	Id           int64    `json:"id"`
	Method       string   `json:"method"`
	SourceLang   string   `json:"source_lang"`
	TargetLang   string   `json:"target_lang"`
}

func TransByDeeplx(src, proxy string, once *sync.Once, wg *sync.WaitGroup, dst chan string) {
	token := os.Getenv("TOKEN")
	if token != "" {
		rep, fail := Deeplx(src)
		if fail != nil {
			return
		} else {
			once.Do(func() {
				fmt.Println("linux.do的DeepLx返回翻译结果")
				dst <- rep
				wg.Done()
			})
		}
	} else {
		//var req DP.DeepLXTranslationResult
		rep, fail := DP.TranslateByDeepLX("auto", "zh", src, "html", proxy, "")
		if fail != nil {
			return
		} else {
			once.Do(func() {
				fmt.Println("DeepLx返回翻译结果", rep)
				dst <- rep.Data
				wg.Done()
			})
		}
	}
}

func TransByGoogle(src, proxy string, once *sync.Once, wg *sync.WaitGroup, dst chan string) {
	cmd := exec.Command("trans", "-brief", "-engine", "google", "-proxy", proxy, ":zh-CN", src)
	output, err := cmd.CombinedOutput()
	result := string(output)
	if err != nil || strings.Contains(string(output), "u001b") || strings.Contains(string(output), "Didyoumean") || strings.Contains(string(output), "Connectiontimedout") {
		log.Printf("google查询命令执行出错\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		return
	} else {
		once.Do(func() {
			fmt.Println("Google返回翻译结果")
			dst <- result
			wg.Done()
		})
	}
}

func TransByBing(src, proxy string, once *sync.Once, wg *sync.WaitGroup, dst chan string) {
	cmd := exec.Command("trans", "-brief", "-engine", "bing", "-proxy", proxy, ":zh-CN", src)
	output, err := cmd.CombinedOutput()
	result := string(output)
	if err != nil || strings.Contains(string(output), "u001b") || strings.Contains(string(output), "Didyoumean") || strings.Contains(string(output), "Connectiontimedout") {
		log.Printf("bing查询命令执行出错\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		return
	} else {
		once.Do(func() {
			fmt.Println("Bing返回翻译结果")
			dst <- result
			wg.Done()
		})
	}
}

func TransOnce(src, proxy string) (string, error) {
	cmd := exec.Command("trans", "-brief", "-engine", "bing", "-proxy", proxy, ":zh-CN", src)
	output, err := cmd.CombinedOutput()
	result := string(output)
	if err != nil || strings.Contains(string(output), "u001b") || strings.Contains(string(output), "Didyoumean") || strings.Contains(string(output), "Connectiontimedout") {
		log.Printf("bing查询命令执行出错\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		return "", err
	}
	return result, nil
}

func Deeplx(src string) (dst string, err error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	data := map[string]string{
		"text":        src,
		"source_lang": "auto",
		"target_lang": "zh",
	}
	token := os.Getenv("TOKEN")
	if token == "" {
		notfound := errors.New("没有找到deeplx的apikey环境变量$TOKEN")
		return "", notfound
	}
	uri := strings.Join([]string{"https://api.deeplx.org", token, "translate"}, "/")
	j, err := util.HttpPostJson(headers, data, uri)
	if err != nil {
		return "deeplx 请求发生错误", err
	}
	fmt.Println(string(j))
	var result DeeplxRep
	json.Unmarshal(j, &result)
	return result.Data, nil
}
