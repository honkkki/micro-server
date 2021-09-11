package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

func main() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			ContextPath: "/nacos",
			Port: 8848,
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "micro-sysconfig",
		Group: "micro",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println(data)

		},
	})

	if err != nil {
		log.Fatal(err)
	}
	for  {
		
	}
}