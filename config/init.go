package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func Init() (config Config, err error) {

	data, err := ioutil.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(data, &config)
	}
	return
}
