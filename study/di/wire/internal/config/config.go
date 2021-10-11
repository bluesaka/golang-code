package config

import (
	"github.com/google/wire"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

// 将NewConfig方法声明为Provider，表示NewConfig方法可以作为一个被别人依赖的对象
var Provider = wire.NewSet(NewConfig)

// NewConfig 初始化配置
func NewConfig() (*Config, error) {
	bt, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(bt, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Config struct {
	Redis RedisCfg `json:"redis"`
}

type RedisCfg struct {
	Addr        string `json:"addr" yaml:"addr"`
	Password    string `json:"password" yaml:"password"`
	DB          int    `json:"db" yaml:"db"`
	MaxActive   int    `json:"max_active" yaml:"max_active"`
	MaxIdle     int    `json:"max_idle" yaml:"max_idle"`
	IdleTimeout int    `json:"idle_timeout" yaml:"idle_timeout"`
	DialTimeout int    `json:"dial_timeout" yaml:"dial_timeout"`
}
