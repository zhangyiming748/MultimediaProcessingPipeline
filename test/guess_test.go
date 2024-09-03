package t

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
	"time"
)

func TestGuess(t *testing.T) {
	//uri:="https://eu.microsub.org/api/v1/client/subscribe?token=660721399cd86c11f92ee83a06428635"
	//prefix:="https://eu.microsub.org/api/v1/client/subscribe?token=660721399cd86c1"
	//suffix:="92ee83a06428635"
	every()
}

func every() {
	// 定义字符集，包括数字和大小写字母
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var count int
	// 遍历所有可能的组合

	for j := 0; j < len(charset); j++ {
		// 生成组合
		combination := string(charset[j])
		// 构造完整的 URL
		uri := fmt.Sprintf("https://eu.microsub.org/api/v1/client/subscribe?token=660721399cd86c1%sf92ee83a06428635", combination)
		// 打印 URL
		fmt.Println(uri)
		if err := HttpGet(uri); err != nil {
			fmt.Printf("url:%v不正确\n", uri)
		} else {
			file, _ := os.OpenFile("subscribe.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
			file.WriteString(uri)
			return
		}
		count++
		time.Sleep(2 * time.Second)
	}

	fmt.Println(count)
}
func TestHttpGet(t *testing.T) {
	err := HttpGet("https://eu.microsub.org/api/v1/client/subscribe?token=660721399cd86c11f92ee83a06428635")
	if err != nil {
		log.Fatal(err)
	}
}
func HttpGet(uri string) error {
	proxyURL, err := url.Parse("http://192.168.1.20:8889") // 替换为你的代理地址和端口
	if err != nil {
		log.Fatalf("Error parsing proxy URL: %v", err)
	}

	// 创建 HTTP 客户端并设置代理
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport}

	// 定义要请求的 URL
	requestURL := uri

	// 创建请求
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// 设置 Chrome 浏览器的请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")

	// 发送请求
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer response.Body.Close()

	// 检查响应状态码
	if response.StatusCode != http.StatusOK {
		e := fmt.Sprintf("Error: received status code %d", response.StatusCode)
		return errors.New(e)
	}

	// 读取响应体
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// 打印响应体
	fmt.Printf("响应体%v\n", string(body))
	return nil
}
