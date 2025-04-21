package translateShell

import (
	"testing"
)

func TestTranslateShell(t *testing.T) {
	src := "hello"
	dst := TransByServer(src)
	t.Log(dst)
}
