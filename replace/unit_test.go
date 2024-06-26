package replace

import (
	"testing"
)

func TestForFileName(t *testing.T) {
	str := "Hello, 世界！123abc!@#$%^&*()_+"
	ret := ForFileName(str)
	t.Log(ret)
}
