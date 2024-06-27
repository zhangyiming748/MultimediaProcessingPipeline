package translateShell

import (
	"Multimedia_Processing_Pipeline/constant"
	"log"
	"os/exec"
	"strings"
)

func TransByGoogle(src string, c *constant.Count, p *constant.Param) (dst string, err error) {
	cmd := exec.Command("trans", "-brief", "-engine", "google", "-proxy", p.GetProxy(), ":zh-CN", src)
	out, err := cmd.CombinedOutput()
	if err != nil || strings.Contains(string(out), "Didyoumean") {
		log.Printf("google查询命令执行出错\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		return "", err
	} else {
		log.Println(dst)
	}
	dst = string(out)
	c.SetGoogle()
	return dst, nil
}
