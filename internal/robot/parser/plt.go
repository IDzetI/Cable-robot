package robot_parser

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

type Plt struct {
	up    float64
	down  float64
	start []float64
}

func (plt *Plt) Read(file string) (trajectory [][]float64, err error) {
	if len(plt.start) != 3 {
		plt.start = []float64{0, 0, 0}
	}

	//read file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")

	//parse
	current := plt.start
	for _, line := range lines {
		var x, y, z float64
		switch line[:2] {
		case "PU":
			z = plt.up
			break
		case "PD":
			z = plt.down
			break
		default:
			continue
		}

		x, y, err = pltLineToXY(line[2:])
		if err != nil {
			return
		}

		if current[2] != z {
			current[2] = z
			trajectory = append(trajectory, current)
		}
		current = []float64{x, y, z}
		trajectory = append(trajectory, current)
	}
	return
}

func pltLineToXY(line string) (x float64, y float64, err error) {
	xy := strings.Split(line, " ")
	if len(xy) != 2 {
		err = errors.New("invalid plt line: " + line)
	}
	x, err = strconv.ParseFloat(xy[0], 64)
	y, err = strconv.ParseFloat(xy[1][:len(xy[1])-2], 64)
	return
}
