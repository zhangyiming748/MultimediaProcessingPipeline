package translateShell

import "fmt"

type Count struct {
	Bing   uint64
	Google uint64
	Deeplx uint64
	Cache  uint64
}

func (c *Count) SetBing() {
	c.Bing++
}

func (c *Count) SetGoogle() {
	c.Google++
}

func (c *Count) SetDeeplx() {
	c.Deeplx++
}

func (c *Count) SetCache() {
	c.Cache++
}

func (c *Count) GetAll() {
	fmt.Printf("从bing获取:%d条\n从google获取:%d条\n从deeplx获取:%d条\n从cache获取:%d条\n", c.Bing, c.Google, c.Deeplx, c.Cache)
}
