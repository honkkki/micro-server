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
	"os"
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
		Type     string `yaml:"type"`
		Ip       string `yaml:"ip"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		MaxIdle  int    `yaml:"max_idle"`
		MaxOpen  int    `yaml:"max_open"`
	} `yaml:"db"`

	Redis *struct {
		Ip   string `yaml:"ip"`
		Port int    `yaml:"port"`
	} `yaml:"redis"`
	Gin *struct {
		RunMode string `yaml:"run_mode"`
	} `yaml:"gin"`
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
	exist, _ := file.Exists(cacheDir)
	if !exist {
		os.Mkdir(cacheDir, 0666)
	}

	// get config
	content, err := NacosClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})

	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(cacheFile, []byte(content), 0666)
	if err != nil {
		log.Fatal(err)
	}

	err = loadFile(cacheFile, nacosConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = NacosClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			err := ioutil.WriteFile(cacheFile, []byte(data), 0666)
			if err != nil {
				log.Println(err)
				return
			}

			err = loadFile(cacheFile, nacosConfig)
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

func loadFile(file string, data interface{}) error {
	err := config.LoadFile(file)
	if err != nil {
		return err
	}

	err = config.Scan(data)
	if err != nil {
		return err
	}

	return nil
}
