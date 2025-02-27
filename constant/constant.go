package constant

import (
	"github.com/fatih/color"
)

const (
	HASEE = "DESKTOP-VGFTVD8"
)

type Param struct {
	VideosLocation string // 视频文件位置
	Language       string // 视频文件语言 English German Russian Japanese Korean Spanish French
	Pattern        string //视频扩展名
	Model          string //whisper 所使用的模型等级 large
	ToolsLocation  string //whisper sensitive.txt 模型保存的位置 如果为空保存在视频文件夹
	Proxy          string // 翻译所需要的网络环境
	Merge          bool   //是否合并字幕 false 不合并 true 合并
	Lines          string // 保存下载url的文档 默认放在root下 文件名为 link.list
	MysqlHost          string // 数据库ip
	MysqlPort          string // 数据库端口
	MysqlUser          string // 数据库用户名
	MysqlPass         string // 数据库密码
	TransService   string // 本地翻译服务位置
}

func (p *Param) GetLines() string {
	return p.Lines
}
func (p *Param) SetLines(s string) {
	p.Lines = s
}
func (p *Param) SetMerge() {
	p.Merge = true
}

func (p *Param) GetMerge() bool {
	return p.Merge
}

func (p *Param) GetVideosLocation() string {
	return p.VideosLocation
}
func (p *Param) SetVideosLocation(r string) {
	p.VideosLocation = r
}

func (p *Param) GetLanguage() string {
	return p.Language
}
func (p *Param) SetLanguage(l string) {
	p.Language = l
}

func (p *Param) GetPattern() string {
	return p.Pattern
}
func (p *Param) SetPattern(pa string) {
	p.Pattern = pa
}

func (p *Param) GetModel() string {
	return p.Model
}

func (p *Param) SetModel(m string) {
	p.Model = m
}
func (p *Param) GetToolsLocation() string {
	if p.ToolsLocation == "" {
		return p.VideosLocation
	}
	return p.ToolsLocation
}

func (p *Param) SetToolsLocation(l string) {
	p.ToolsLocation = l
}

func (p *Param) GetProxy() string {
	return p.Proxy
}

func (p *Param) SetProxy(pr string) {
	p.Proxy = pr
}
func (p *Param) SetMysql() {
	p.MysqlUser = "root"
	p.MysqlPass= "163453"
	p.MysqlHost = "192.168.2.8"
	p.MysqlUser = "3306"
}


func (p *Param) SetTransService(s string) {
	p.TransService = s
}

func (p *Param) GetTransService() string {
	return p.TransService
}

type Count struct {
	Bing   uint64
	Google uint64
	Deeplx uint64
	Cache  uint64
}

func (c *Count) SetBing() {
	c.Bing++
}

func (c *Count) GetBing() uint64 {
	return c.Bing
}

func (c *Count) SetGoogle() {
	c.Google++
}

func (c *Count) GetGoogle() uint64 {
	return c.Google
}

func (c *Count) SetDeeplx() {
	c.Deeplx++
}

func (c *Count) GetDeeplx() uint64 {
	return c.Deeplx
}

func (c *Count) SetCache() {
	c.Cache++
}

func (c *Count) GetCache() uint64 {
	return c.Cache
}

func Warning(s string) {
	c := color.New()
	c.Add(color.FgYellow)
	c.Add(color.BgWhite)
	c.Add(color.Bold)
	c.Println(s)
}
func Error(s string) {
	c := color.New()
	c.Add(color.FgRed)
	c.Add(color.BgWhite)
	c.Add(color.Bold)
	c.Println(s)
}
func Info(s string) {
	c := color.New()
	c.Add(color.FgHiBlack)
	c.Add(color.BgWhite)
	c.Add(color.Bold)
	c.Println(s)
}
