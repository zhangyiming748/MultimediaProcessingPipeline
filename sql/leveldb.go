package sql

import (
	"Multimedia_Processing_Pipeline/constant"
	"log"
	"os"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

var levelDB *leveldb.DB

func SetLevelDB(p *constant.Param) {
	location := strings.Join([]string{p.GetVideosLocation(), "leveldb"}, string(os.PathSeparator))
	db, err := leveldb.OpenFile(location, nil)
	if err != nil {
		log.Fatalf("leveldb数据库创建失败:%v\n", err)
	}
	levelDB = db
}
func GetLevelDB() *leveldb.DB {
	return levelDB
}
