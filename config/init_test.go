package config

import (
	"reflect"
	"testing"
)

func TestInitConfig(t *testing.T) {
	InitConfig()
	t.Log(appConfig.Config.Address)
	t.Log(appConfig.Config.Port)
	t.Log(reflect.TypeOf(appConfig.Config.Port))


}
