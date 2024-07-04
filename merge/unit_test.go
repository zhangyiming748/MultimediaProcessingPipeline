package merge

import (
	"Multimedia_Processing_Pipeline/constant"
	mylog "Multimedia_Processing_Pipeline/log"
	"testing"
)

func init() {

}

// go test -v -run TestMerge
func TestMerge(t *testing.T) {
	p := new(constant.Param)
	p.Root = "/mnt/c/Users/zen/Downloads"
	p.Language = "Japanese"
	p.Pattern = "mp4"
	p.Model = "small"
	p.Location = "/mnt/c/Users/zen/Downloads"
	p.Proxy = "192.168.1.20:8889"
	p.Merge = false
	mylog.SetLog(p)
	MkvWithAss("/mnt/d/merge/sdde-712.mp4", p)
}
