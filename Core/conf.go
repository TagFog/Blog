package Core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"web/Config"
	"web/Global"
)

// 读取配置
func InitConf() {
	const ConfigFile = "./setting.yaml"
	c := &Config.Config{}
	yamlConfig, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("Get yamlConfig error: %s", err))
	}

	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		log.Fatalf("Config init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success")
	Global.Config = c
}
