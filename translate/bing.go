package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	"log"
	"os/exec"
	"strings"
)

func TransByBing(proxy, language, src string, c *constant.Count) {
	cmd := exec.Command("trans", "-brief", "-engine", "bing", "-proxy", proxy, language, src)
	output, err := cmd.CombinedOutput()
	if err != nil || strings.Contains(string(output), "u001b") || strings.Contains(string(output), "Didyoumean") || strings.Contains(string(output), "Connectiontimedout") {
		log.Printf("bing查询命令执行出错\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		return
	}

	c.SetBing()

}
