package config

import (
	"flag"
	"github.com/spf13/viper"
)

var (
	configPath = flag.String("f", "./config.yaml", "the config file")
)

func init() {
	flag.Parse()
	viper.SetConfigFile(*configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
