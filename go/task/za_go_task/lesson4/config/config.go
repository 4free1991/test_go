package config

import (
	"app/lesson4/pkg/logger"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type config struct {
	AppName string `yaml:"appName"`
	Stats   struct {
		Brokers []string `yaml:"brokers"`
		Topics  []string `yaml:"topics"`
		GroupID string   `yaml:"groupId"`
	} `yaml:"stats"`
	Redis struct {
		Addr         string `yaml:"addr"`
		Password     string `yaml:"password"`
		Db           int    `yaml:"db"`
		DialTimeout  int    `yaml:"dialTimeout"`
		WriteTimeout int    `yaml:"writeTimeout"`
		ReadTimeout  int    `yaml:"readTimeout"`
	} `yaml:"redis"`
	Mongodb struct {
		URL    string `yaml:"url"`
		Dbname string `yaml:"dbname"`
	} `yaml:"mongodb"`
}

var _config *config

func GetConfig() *config {
	if _config == nil {
		panic("config is null")
	}
	return _config
}

func InitConfig() {
	yamlFile, err := os.ReadFile("./lesson4/config/config.yml")
	if err != nil {
		panic(err)
	}

	_config = &config{}

	err = yaml.Unmarshal(yamlFile, &_config)
	if err != nil {
		fmt.Println(err.Error())
	}
	logger.Init()

}
