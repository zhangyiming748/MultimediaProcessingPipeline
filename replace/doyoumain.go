package replace

import (
	"strings"
)

func DoYouMean(src string) (dst string) {
	if strings.Contains(src, "Didyoumean") {
		sp := strings.Split(src, "[0m")
		dst = sp[1]
	}
	return dst
}

/*
[33mDidyoumean[1mаможетдажеприсомнойувами?Гдевырассказываетеотом,кактрохнадобыменя?Пожалуйста,[22m[0m
*/
