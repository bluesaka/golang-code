package config

import (
	"fmt"
	"github.com/ghodss/yaml"
	jsoniter "github.com/json-iterator/go"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/cast"
	"log"
	"os"
	"strings"
)

const (
	EnvKey      = "NACOS_HOST"
	NamespaceId = "c00aeff2-561b-4ba4-91cd-06da2ef5eb6b"
	Group       = "sgs"
	FileType    = "yaml"
)

var dataIds = []string{"sgs-test"}

type nacos struct {
	host        string
	namespaceId string
	client      config_client.IConfigClient
}

func InitNacos() {
	nacosHostEnv := os.Getenv(EnvKey)
	nacosHosts := strings.Split(nacosHostEnv, ":")
	if len(nacosHosts) < 2 {
		panic("nocas host env err:" + nacosHostEnv)
	}

	serverCfg := []constant.ServerConfig{
		{
			IpAddr: nacosHosts[0],
			Port:   cast.ToUint64(nacosHosts[1]),
		},
	}

	clientCfg := &constant.ClientConfig{
		NamespaceId: NamespaceId,
		TimeoutMs:   5000,
		LogLevel:    "error",
		//NotLoadCacheAtStart: true,
		//LogDir:              "/tmp/nacos/log",
		//CacheDir:            "/tmp/nacos/cache",
		//RotateTime:          "1h",
		//MaxAge:              3,
	}

	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  clientCfg,
		ServerConfigs: serverCfg,
	})
	if err != nil {
		panic(err)
	}

	ns := nacos{
		client: client,
	}

	for _, dataId := range dataIds {
		ns.setConfig(dataId, Group)
		go ns.listenConfig(dataId, Group)
	}
}

func (n *nacos) setConfig(dataId, group string) {
	content, err := n.client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})

	if err != nil {
		panic(fmt.Sprintf("nacos failed to set config from dataId: %s, group: %s, err: %v\n", dataId, group, err))
	}

	switch FileType {
	case "yaml":
		setConfigWithYaml(content)
	case "json":
		setConfigWithJson(content)
	default:
		panic(fmt.Sprintf("nacos not support file format %s\n", FileType))
	}
}

func (n *nacos) listenConfig(dataId, group string) {
	err := n.client.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			n.setConfig(dataId, group)
		},
	})

	if err != nil {
		log.Println("nacos listenConfig error:", err)
	}
}

func setConfigWithYaml(content string) {
	var m map[string]interface{}
	if err := yaml.Unmarshal([]byte(content), &m); err != nil {
		log.Println("nacos setConfigWithYaml error: ", err)
	}
	fmt.Printf("config: %+v\n", m)
}

func setConfigWithJson(content string) {
	var m map[string]interface{}
	if err := jsoniter.Unmarshal([]byte(content), &m); err != nil {
		log.Println("nacos setConfigWithJson error: ", err)
	}
	fmt.Printf("config: %+v\n", m)
}
