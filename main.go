package main

import (
	"Multimedia_Processing_Pipeline/constant"
	"Multimedia_Processing_Pipeline/log"
	"Multimedia_Processing_Pipeline/sql"
)

func initConfig(p *constant.Param) {
	log.SetLog(p)
	sql.SetDatabase(p)
}
func main() {

}
