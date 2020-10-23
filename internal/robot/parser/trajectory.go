package robot_parser

import (
	"errors"
	"github.com/IDzetI/Cable-robot/pkg/utils"
	"io/ioutil"
	"strconv"
	"strings"
)

type Rt struct{}

func (rt *Rt) Read(file string, c chan string) (trajectory [][]float64, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		strValues := strings.Split(line, " ")
		if len(strValues) != 3 {
			err = errors.New("file corrupted: line " + strconv.Itoa(i) + " (" + line + ")")
			return
		}
		var xyz []float64
		for _, coordinate := range strValues {
			var value float64
			value, err = strconv.ParseFloat(coordinate, 64)
			if err != nil {
				return
			}
			xyz = append(xyz, value)
		}
		trajectory = append(trajectory, xyz)
		c <- "OK " + utils.ToString(xyz)
	}
	c <- "load complete"
	return
}
