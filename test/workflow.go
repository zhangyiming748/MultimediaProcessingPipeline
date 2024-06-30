package t

import (
	"fmt"
	"sync"
)

func main() {
	videos := []string{
		"视频0",
		"视频1",
		"视频2",
		"视频3",
		"视频4",
		"视频5",
		"视频6",
		"视频7",
		"视频8",
		"视频9",
		"视频10",
		"视频11",
		"视频12",
		"视频13",
		"视频14",
		"视频15"}

	// Create channels for communication between steps
	step1Ch := make(chan string)
	step2Ch := make(chan string)
	step3Ch := make(chan string)
	step4Ch := make(chan string)

	// Create a wait group to track concurrent tasks
	var wg sync.WaitGroup

	// Launch a goroutine for each step
	go func() {
		for v := range step1Ch {
			wg.Add(1)
			go func(v string) {
				defer wg.Done()
				uuid := step2(v)
				step2Ch <- uuid
			}(v)
		}
	}()

	go func() {
		for uuid := range step2Ch {
			wg.Add(1)
			go func(uuid string) {
				defer wg.Done()
				step3(uuid)
				step3Ch <- uuid
			}(uuid)
		}
	}()

	go func() {
		for uuid := range step3Ch {
			wg.Add(1)
			go func(uuid string) {
				defer wg.Done()
				step4(uuid)
				step4Ch <- uuid
			}(uuid)
		}
	}()

	// Process videos in a pipeline fashion
	for _, v := range videos {
		fileName := step1(v)
		step1Ch <- fileName
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close channels to signal completion
	close(step1Ch)
	close(step2Ch)
	close(step3Ch)
	close(step4Ch)
}

/*
模拟下载函数
*/
func step1(s string) (fileName string) {
	fmt.Println("下载视频")
	fileName = "视频" + s
	return fileName
}

/*
模拟生成字幕函数
*/
func step2(fileName string) string {
	fmt.Println("视频根据文件名生成字幕")
	return fileName + "_字幕"
}

/*
模拟翻译字幕函数
*/
func step3(subtitle string) string {
	fmt.Println("视频根据字幕翻译字幕")
	return subtitle + "_翻译"
}

/*
模拟合并视频函数
*/
func step4(translatedSubtitle string) {
	fmt.Println("视频根据翻译字幕合成字幕")
	fmt.Println("合成后的字幕文件名为:", translatedSubtitle)
}
