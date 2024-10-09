package translateShell

import (
	"fmt"
	DeepLx "github.com/zhangyiming748/DeepLX"
	"log"
	"os/exec"
	"strings"
	"sync"
)

func TransByDeeplx(src, proxy string, once *sync.Once, wg *sync.WaitGroup, dst chan string) {
	result, fail := DeepLx.TranslateByDeepLX("auto", "zh", src, proxy)
	if fail != nil {
		return
	} else {
		once.Do(func() {
			fmt.Println("DeepLx返回翻译结果")
			dst <- result
			wg.Done()
		})
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
