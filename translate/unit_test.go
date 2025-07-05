package translateShell

import (
	"testing"
)

func TestTranslateShell(t *testing.T) {
	src := "hello"
	dst := TransByServer(src)
	t.Logf("翻译结果: %s", dst)
	t.Log(dst)
}
