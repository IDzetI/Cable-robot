package robot_parser

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

type Rt struct{}

func (rt *Rt) Read(file string) (trajectory [][]float64, err error) {
	//read file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
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
	}
	return
}
