package fin

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

var (
	configPath  string
	messagePath string
	logPath     string
)

func setConfig(path string) {
	configPath = path
}

func setMessage(path string) {
	messagePath = path
}

func setLog(path string) {
	logPath = path
}

func Config(key string) interface{} {
	fileData, _ := ioutil.ReadFile(configPath)
	var config map[interface{}]interface{}
	yaml.Unmarshal(fileData, &config)
	config = config[Mode()].(map[interface{}]interface{})
	keys := strings.Split(key, ".")
	length := len(keys)
	if length == 1 {
		return config[key]
	} else {
		var value interface{}
		for _, v := range keys {
			if value == nil {
				value = config[v]
			} else {
				value = value.(map[interface{}]interface{})[v]
			}
		}
		return value
	}
}

func Message(status int) string {
	value := getMessage(status, messagePath)
	if value == "" {
		return baseMessage(status)
	}
	return value
}

func baseMessage(status int) string {
	value := getMessage(status, "./message.yml")
	if value == "" {
		return baseMessage(11000)
	}
	return value
}

func getMessage(status int, path string) string {
	var msg map[int]string
	fileData, _ := ioutil.ReadFile(path)
	yaml.Unmarshal(fileData, &msg)
	return msg[status]
}
