package translateShell

import (
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
	result, err := Req(src)
	if err != nil {
		TransByServer(src)
	}
	result = strings.Replace(result, "\\r\\n", "", 1)
	result = strings.Replace(result, "\n", "", 1)
	result = strings.Replace(result, "\r\n", "", 1)
	if result == "" {
		return
	}
	return result
}