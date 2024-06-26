package constant

type Param struct {
	Root     string // 视频文件位置
	Language string // 视频文件语言 English German Russian Japanese Korean Spanish French
	Pattern  string //视频扩展名
	Model    string //whisper 所使用的模型等级 large
	Location string //whisper 模型保存的位置 如果为空保存在视频文件夹
	Proxy    string // 翻译所需要的网络环境
}

func (p *Param) GetRoot() string {
	return p.Root
}
func (p *Param) SetRoot(r string) {
	p.Root = r
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
func (p *Param) GetLocation() string {
	if p.Location == "" {
		return p.Root
	}
	return p.Location
}
func (p *Param) SetLocation(l string) {
	p.Location = l
}
func (p *Param) GetProxy() string {
	return p.Proxy
}
func (p *Param) SetProxy(pr string) {
	p.Proxy = pr
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
