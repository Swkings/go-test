package test

import (
	"fmt"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Name     string
	ListenOn string
	Timeout  int64
}

func TestConfig(t *testing.T) {
	var configFile string = "etc/config.yaml"
	var c Config
	conf.MustLoad(configFile, &c)
	fmt.Printf("config: %v\n", PrettyMapStruct(c, true))
}
