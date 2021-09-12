package config

import (
	"reflect"
	"testing"
)

func TestInitConfig(t *testing.T) {
	InitConfig()
	t.Log(ConfigData.Config.Address)
	t.Log(ConfigData.Config.Port)
	t.Log(reflect.TypeOf(ConfigData.Config.Port))


}
