package utils

import (
	"fmt"
	"io/ioutil"
)

//read all file
func ReadAll(file string) string {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	return string(s)
}
