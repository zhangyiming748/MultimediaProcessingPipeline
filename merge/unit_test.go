package merge

import (
	"Multimedia_Processing_Pipeline/constant"
	"testing"
)

func init() {

}

// go test -v -run TestMerge
func TestMerge(t *testing.T) {
	p := new(constant.Param)
	MkvWithAss("/media/zen/swap/telegram.mp4", p)
}
