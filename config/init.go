package config

import (
	"fmt"
	"github.com/honkkki/micro-server/tools"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/file"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	Config *struct {
		Address string
		Path    string
		Port    uint64
	}

	*DataConfig
}

type DataConfig struct {
	Db *struct {
		Ip   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"db"`

	Redis *struct {
		Ip   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"redis"`
}

var ConfigData *AppConfig
var NacosClient config_client.IConfigClient

func InitConfig() {
	configFile := tools.GetRootDir() + "/app.yaml"
	err := config.LoadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	ConfigData = &AppConfig{
		Config:     nil,
		DataConfig: new(DataConfig),
	}

	err = config.Get("micro").Scan(ConfigData)
	if err != nil {
		log.Fatal(err)
	}

	// 创建nacos客户端
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      ConfigData.Config.Address,
			ContextPath: ConfigData.Config.Path,
			Port:        ConfigData.Config.Port,
		},
	}

	NacosClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})

	if err != nil {
		log.Fatal(err)
	}

	watchConfig("micro-sysconfig", "micro", ConfigData.DataConfig)
}

func watchConfig(dataId, group string, nacosConfig interface{}) {
	cacheDir := fmt.Sprintf("%s/config/cache/", tools.GetRootDir())
	cacheFileName := fmt.Sprintf("%s_%s.yaml", group, dataId)
	cacheFile := cacheDir + cacheFileName

	exist, _ := file.Exists(cacheFile)
	if !exist {
		// get config
	}

	err := NacosClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {

			err := ioutil.WriteFile(cacheFile, []byte(data), 0666)
			if err != nil {
				log.Println(err)
				return
			}

			err = config.LoadFile(cacheFile)
			if err != nil {
				log.Println(err)
				return
			}
			err = config.Scan(nacosConfig)
			if err != nil {
				log.Println(err)
				return
			}
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
