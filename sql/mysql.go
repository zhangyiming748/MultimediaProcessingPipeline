package sql

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func SetMysql() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@tcp(192.168.1.9:3306)/Translate?charset=utf8")
	if err != nil {
		panic(err)
	}
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
}
func GetMysql() *xorm.Engine {
	return engine
}
