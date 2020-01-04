package robot

import (
	"errors"
	"github.com/IDzetI/Cable-robot/internal/robot/parser"
	"github.com/IDzetI/Cable-robot/pkg/utils"
	"strconv"
	"strings"
)

type file struct {
	plt        *robot_parser.PltConfig
	trajectory [][]float64
}

func (uc *UseCase) FileLoadPLT(file string, c chan string) (err error) {

	//check config
	if uc.file.plt == nil {
		return errors.New("plt config is empty")
	}

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

func (uc *UseCase) FileLoad(file string, c chan string) (err error) {
	return
}

func (uc *UseCase) FileInit(c chan string) (err error) {
	return
}

func (uc *UseCase) FileNext(c chan string) (err error) {
	return
}

func (uc *UseCase) FileCurrent(c chan string) (err error) {
	return
}

func (uc *UseCase) FilePrevious(c chan string) (err error) {
	return
}

func (uc *UseCase) FileSetCursor(i int, c chan string) (err error) {
	return
}

func (uc *UseCase) FileGo(c chan string) (err error) {
	return
}

func (uc *UseCase) FileRun(c chan string) (err error) {
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
