package log

import (
	"github.com/zhangyiming748/lumberjack"
	"io"
	"log"
	"os"
)

func SetLog() {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   "MultimediaProcessingPipeline.log",
		MaxSize:    1, // MB
		MaxBackups: 1,
		MaxAge:     28, // days
		LocalTime:  true,
	}
	fileLogger.Rotate()
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(3 | 16)
}
