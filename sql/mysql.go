package sql

import (
	"Multimedia_Processing_Pipeline/constant"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func SetMysql(p *constant.Param) {
	var err error
	session := strings.Join([]string{"root:123456@tcp(", p.GetMysql(), ")/Translate?charset=utf8"}, "")
	engine, err = xorm.NewEngine("mysql", session)
	if err != nil {
		panic(err)
	}
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
}
func GetMysql() *xorm.Engine {
	return engine
}
