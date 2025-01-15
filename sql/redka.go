package sql

import (
	"Multimedia_Processing_Pipeline/constant"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nalgeon/redka"
)

var (
	red *redka.DB
)

func GetRedka() *redka.DB {
	return red
}
func SetRedka(p *constant.Param) {
	location := strings.Join([]string{p.GetVideosLocation(), "trans.db"}, string(os.PathSeparator))
	red, _ = redka.Open(location, nil)
}
