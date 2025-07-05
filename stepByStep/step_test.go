package stepbystep

import (
	"testing"
)


//go test -v -timeout 10h -run TestDownloadAll
func TestDownloadAll(t *testing.T){
	file:="/app/stepByStep/links.txt"
	proxy:="127.0.0.1:8889"
	location:="/app/stepByStep"
	links:=ReadLinkToSlice(file)
	for _,link:=range links{
		RunYtdlp(link,proxy,location)
	}
}