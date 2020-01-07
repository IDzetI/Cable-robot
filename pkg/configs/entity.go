package configs

import (
	"io/ioutil"
	"strings"
)

type Config struct {
	config map[string]string
}

func LoadConfig(file string) (config Config, err error) {
	err = config.load(file)
	return
}

func (c *Config) load(file string) (err error) {
	//read file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), configSeparator)
	for _, line := range lines {
		keyValue := strings.Split(line, keyValueSeparator)
		if len(keyValue) < 2 {
			continue
		}
		c.config[keyValue[0]] = line[len(keyValue[0])+1:]
	}
	return
}
