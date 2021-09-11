package main

import (
	"fmt"
	"github.com/honkkki/micro-server/config"
	"time"
)

func main() {
	config.InitConfig()
	fmt.Println(config.ConfigData.DataConfig.Db)

	for  {
		time.Sleep(time.Second)
		fmt.Println(config.ConfigData.DataConfig.Db)
	}

}