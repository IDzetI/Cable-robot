package robot_parser

import (
	"github.com/IDzetI/Cable-robot/pkg/utils"
	"strings"
)

type PltConfig struct {
	up    float64
	down  float64
	start []float64
}

func PLT(file string) (trajectory [][]float64, err error) {
	if len(uc.file.plt.start) != 3 {
		uc.file.plt.start = []float64{0, 0, 0}
	}

	//read file
	lines := strings.Split(utils.ReadAll(file), "\n")

	//parse
	current := uc.file.plt.start
	var trajectory [][]float64
	for _, line := range lines {
		var x, y, z float64
		switch line[:2] {
		case "PU":
			z = uc.file.plt.up
			break
		case "PD":
			z = uc.file.plt.down
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

	uc.file.trajectory = trajectory

	return
}
