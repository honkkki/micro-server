package main

import (
	"fmt"
	"github.com/honkkki/micro-server/config"
	"github.com/micro/go-micro/v2/logger"
	"time"
)

func main() {
	config.InitConfig()

	for  {
		fmt.Println(config.ConfigData.DataConfig.Db)
		logger.Info("get config...")
		time.Sleep(time.Second)
	}

}