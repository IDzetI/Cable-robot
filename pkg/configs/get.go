package configs

import (
	"errors"
	"strconv"
)

func (c *Config) GetString(key string) (s string, err error) {
	s, ok := c.config[key]
	if !ok {
		err = errors.New(keyDoesNotExist + key)
	}
	return
}

func (c *Config) GetFloat64(key string) (v float64, err error) {
	s, err := c.GetString(key)
	if err != nil {
		return
	}
	return strconv.ParseFloat(s, 64)
}
