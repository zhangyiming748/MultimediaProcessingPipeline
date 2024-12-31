package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/util"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"sync"
	"time"

	translate "github.com/OwO-Network/DeepLX/translate"
)

type RaspiRep struct {
	Src    string `json:"src"`
	Dst    string `json:"dst"`
	Source string `json:"source"`
	Target string `json:"target"`
	From   string `json:"from"`
}

func TransByDeeplx(src string, p *constant.Param, once *sync.Once, wg *sync.WaitGroup, dst chan string) {
	rep, fail := Deeplx(src, p)
	if fail != nil {
		return
	} else {
		once.Do(func() {
			fmt.Println("linux.do的DeepLx返回翻译结果")
			dst <- rep
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
func Deeplx(src string, p *constant.Param) (dst string, err error) {
	if checkIPPort(p.GetTransService()) {
		return LinuxdoDeeplx(src, p)
	} else {
		return GithubDeepLx(src)
	}
}

/*
使用linuxdo的deeplx
*/
func LinuxdoDeeplx(src string, p *constant.Param) (dst string, err error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	data := map[string]string{
		"src":         src,
		"source_lang": "auto",
		"target_lang": "zh",
	}
	prefix := p.GetTransService()
	host := strings.Join([]string{prefix, "api/v1/translate"}, "/")
	j, err := util.HttpPostJson(headers, data, host)
	if err != nil {
		return "deeplx 请求发生错误", err
	}
	fmt.Println(string(j))
	var result RaspiRep
	json.Unmarshal(j, &result)
	return result.Dst, nil
}

/*
使用开源deeplx
*/
func GithubDeepLx(src string) (string, error) {
	if rep, err := translate.TranslateByDeepLX("auto", "zh", src, "html", "", ""); err != nil {
		return "", err
	} else {
		return rep.Data, err
	}
}

/*
 */
func checkIPPort(ipPort string) bool {
	conn, err := net.DialTimeout("tcp", ipPort, 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
