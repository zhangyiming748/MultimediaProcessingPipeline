package sql

import (
	"Multimedia_Processing_Pipeline/constant"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nalgeon/redka"
	"os"
	"strings"
)

var (
	db *redka.DB
)

func GetDatabase() *redka.DB {
	return db
}
func SetDatabase(p *constant.Param) {
	location := strings.Join([]string{p.GetRoot(), "trans.db"}, string(os.PathSeparator))
	db, _ = redka.Open(location, nil)
}
