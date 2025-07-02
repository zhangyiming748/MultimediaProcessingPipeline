package translateShell

import (
	"testing"
)

func TestTranslateShell(t *testing.T) {
	src := "hello"
	dst := TransByServer(src,"DrkwqR4tE3DRyOseVibFah62BJXmcIryt4I9rTtzXTs")
	t.Logf("翻译结果: %s", dst)
	t.Log(dst)
}
