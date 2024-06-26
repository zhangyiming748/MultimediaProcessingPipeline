package util

import "testing"

func TestGetQuotedContent(t *testing.T) {
	input := `[FixupM3u8] Fixing MPEG-TS in MP4 container of "Follow my.mp4"`
	content, err := getQuotedContent(input)
	if err != nil {
		return
	} else {
		t.Log(content)
	}
}
