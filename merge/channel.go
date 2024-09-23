package merge

import (
	"fmt"
	"log"
	"sync"
)

func MergeByChannel(file chan string, wg *sync.WaitGroup) {
	// 启动一个 goroutine 来处理通道中的消息
	go func() {
		for {
			// 等待通道中的消息
			fp, ok := <-file
			// 处理消息
			log.Printf("\033[31m通道接收到任务\033[0m\n")
			wg.Add(1)
			Mp4WithSrt(fp)
			wg.Done()
			//fmt.Println("处理消息:", msg)
			if !ok {
				// 如果通道已关闭，退出循环
				fmt.Println("通道已关闭，退出处理")
				return
			}
		}
	}()
}
