package model

import (
	"Multimedia_Processing_Pipeline/sql"
	"time"
)

type YtdlpHistory struct {
	Id        int64     `xorm:"pk autoincr notnull comment('主键id') INT(11)"`
	Url       string    `xorm:"varchar(255) comment(下载网址)"`
	Title     string    `xorm:"varchar(255) comment(标题)"`
	Host      string    `xorm:"varchar(255) comment(从哪台设备下载的)"`
	Source    string    `xorm:"varchar(255) comment(来源 如 ph xvideo)"`
	Key       string    `xorm:"varchar(255) comment(可以被识别为相同视频的关键部分 比如viewkey)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (y *YtdlpHistory) InsertOne() (int64, error) {
	return sql.GetMysql().InsertOne(y)
}

func (y *YtdlpHistory) FindByUrl() (bool, error) {
	return sql.GetMysql().Where("url = ?", y.Url).Get(y)
}

func (y *YtdlpHistory) FindByTitle() (bool, error) {
	return sql.GetMysql().Where("title = ?", y.Title).Get(y)
}

func (y *YtdlpHistory) InsertAll(ytdlpHistory []YtdlpHistory) (int64, error) {
	return sql.GetMysql().Insert(ytdlpHistory)
}
