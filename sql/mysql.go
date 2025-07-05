package sql

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func SetMysql(user, password, host, port, dbname string) {

	var err error
	//user := p.MysqlUser
	//password := p.MysqlPass
	//host := p.MysqlHost
	//port := p.MysqlPort
	//dbname:="Translate"
	// 先连接到 MySQL 服务器（不指定数据库）
	rootDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", user, password, host, port)
	tempEngine, err := xorm.NewEngine("mysql", rootDSN)
	if err != nil {
		log.Printf("连接MySQL服务器失败: %v\n", err)

		return
	}

	// 检查数据库是否存在
	rows, err := tempEngine.QueryString("SELECT SCHEMA_NAME FROM information_schema.SCHEMATA WHERE SCHEMA_NAME = 'Translate'")
	if err != nil {
		log.Printf("查询数据库失败: %v\n", err)

		return
	}

	// 如果数据库不存在，创建它
	if len(rows) == 0 {
		_, err = tempEngine.Exec("CREATE DATABASE `Translate` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci'")
		if err != nil {
			log.Printf("创建数据库失败: %v\n", err)
			return
		}
		log.Println("成功创建数据库 tdl")
	}

	// 关闭临时连接
	tempEngine.Close()

	// 连接到 tdl 数据库
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4", user, password, host, port, dbname)
	engine, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		log.Printf("连接tdl数据库失败: %v\n", err)
		return
	}

	if err = engine.Ping(); err != nil {
		log.Printf("连接数据库失败: %v\n", err)
		return
	} else {
		log.Printf("成功Ping到数据库\n")
		engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	}

	log.Printf("成功连接到数据库\n")
}
func GetMysql() *xorm.Engine {
	return engine
}
