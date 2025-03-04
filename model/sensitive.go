package model

import (
	"Multimedia_Processing_Pipeline/sql"
	"time"
)

type Sensitive struct {
	Id        int64     `xorm:"pk autoincr notnull comment('主键id') INT(11)"`
	Before    string    `xorm:"varchar(255) comment(替换前)"`
	After     string    `xorm:"varchar(255) comment(替换后)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (s *Sensitive) InsertOne() (int64, error) {
	return sql.GetMysql().InsertOne(s)
}

func (s *Sensitive) FindBySrc() (bool, error) {
	return sql.GetMysql().Where("before = ?", s.Before).Get(s)
}

func (s *Sensitive) GetAll() (ss []Sensitive) {
	err := sql.GetMysql().Find(&ss)
	if err != nil {
		return nil
	}
	return ss
}
