package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var ExitAfterDone = false

func GetExitStatus() bool {
	return ExitAfterDone
}

func SetExitStatus(b bool) {
	log.Println("改变退出状态")
	ExitAfterDone = b
}
func ExitAfterRun() {
	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			input, _ := reader.ReadString('\n')
			fmt.Printf("You entered is %T\t%v", input, input)
			if input == "q\n" || input == "q\r\n" {
				log.Println("接收到q")
				SetExitStatus(true)
				log.Printf("退出状态改变,新值=%v\n", ExitAfterDone)
			}
		}
	}()
}
