package replace

import "strings"

/*
终端输出内容转换为可以保存到mysql的字段
*/
func TTY2Mysql(tty string) string {
	tty = strings.Replace(tty, "  ", " ", -1)
	tty = strings.Replace(tty, "\u0000", " ", -1)
	tty = strings.Replace(tty, "\n", " ", -1)
	return tty
}
