package model

import (
	"Multimedia_Processing_Pipeline/sql"
	"time"
)

type TranslateHistory struct {
	Id  int64  `xorm:"comment('主键id') INT(11)"`
	Src string `xorm:"varchar(255) comment(原文)"`
	Dst string `xorm:"varchar(255) comment(译文)"`
	//Source_lang string    `xorm:"varchar(255) comment(源语言)"`
	//Target_lang string    `xorm:"varchar(255) comment(目标语言)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (t TranslateHistory) InsertOne() (int64, error) {
	return sql.GetMysql().InsertOne(t)
}
func (t TranslateHistory) InsertAll(histories []TranslateHistory) (int64, error) {
	return sql.GetMysql().Insert(histories)
}
