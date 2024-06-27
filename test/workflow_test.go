package test

import (
	"fmt"
	"testing"
)

func TestMaster(t *testing.T) {
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
	limit := make(chan struct{}, 1)
	download := make(chan string, 1)
	step1(download, limit)
	//subtitle := make(chan string, 1)
	//step2(subtitle, limit)
	//trans := make(chan string, 1)
	//step3(trans, limit)
	//merge := make(chan string, 1)
	//step4(merge, limit)
	for _, uri := range videos {
		download <- uri
	}
}

/*
模拟下载函数
1 传进来一个网址
2 程序开始干活
3 输出一个下载完成的信号
*/
func step1(ch chan string, limit chan struct{}) {
	for {
		mission := <-ch
		limit <- struct{}{}
		go func() {
			timer(5, "下载中")
			fmt.Println(mission, "下载完成")
			<-limit
		}()
		<-limit
	}
}

//
///*
//模拟生成字幕函数
//*/
//func step2(ch chan string, limit chan struct{}) {
//
//}
//
///*
//模拟翻译字幕函数
//*/
//func step3(ch chan string, limit chan struct{}) {
//
//}
//
///*
//模拟合并视频函数
//*/
//func step4(ch chan string, limit chan struct{}) {
//
//}

func timer(t int, msg string) {
	for i := t; i > 0; i-- {
		fmt.Println(msg)
	}
}
