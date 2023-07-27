package main

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

// Configuration 项目配置
type Configuration struct {
	// gtp apikey
	ApiKey string `yaml:"api_key"`
	// 使用模型
	Model string `yaml:"model"`
	// 超时时间
	SessionTimeout int `yaml:"session_timeout"`
	// 单人单日请求次数上限，默认为0，即不限制
	MaxRequest int `yaml:"max_request"`
	// 指定服务启动端口，默认为 8090
	Port string `yaml:"port"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{}
		data, err := ioutil.ReadFile("config.yml")
		if err != nil {
			log.Fatal(err)
		}
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatal("sssssssssssssss", err)
		}

	})
	return config
}
