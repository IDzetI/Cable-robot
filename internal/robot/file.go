package robot

import (
	"errors"
	"github.com/IDzetI/Cable-robot/internal/robot/parser"
	"strings"
)

type file struct {
	trajectory [][]float64
	cursor     int
	plt        *robot_parser.Plt
	tr         robot_parser.Rt
}

func (uc *UseCase) ConfigPLT(up, down float64, start []float64) (err error) {
	if len(start) != 3 {
		return errors.New("invalid start point")
	}
	uc.file.plt = &robot_parser.Plt{
		Up:    up,
		Down:  down,
		Start: start,
	}
	return
}

func (uc *UseCase) FileLoad(fileString string, c chan string) (err error) {

	uc.file = &file{}

	fileName := strings.Split(fileString, ".")

	var trajectory [][]float64

	switch fileName[len(fileName)-1] {
	case "plt":
		if uc.file.plt == nil {
			err = errors.New("plt config is empty")
			return
		}
		trajectory, err = uc.file.plt.Read(fileString, c)

	case "rt":
		trajectory, err = uc.file.tr.Read(fileString, c)

	default:
		err = errors.New("incorrect file extension")
		return
	}

	if err != nil {
		return
	}
	uc.file.trajectory = trajectory
	uc.file.cursor = 0
	return
}

func (uc *UseCase) FileInit(c chan string) (err error) {
	return uc.MoveInJoinSpace(uc.file.trajectory[uc.file.cursor], c)
}

func (uc *UseCase) FileNext(c chan string) (err error) {
	err = uc.FileSetCursor((uc.file.cursor+1)%len(uc.file.trajectory), c)
	if err != nil {
		return
	}
	return uc.FileCurrent(c)
}

func (uc *UseCase) FileCurrent(c chan string) (err error) {
	return uc.MoveInCartesianSpace(uc.file.trajectory[uc.file.cursor], c)
}

func (uc *UseCase) FilePrevious(c chan string) (err error) {
	err = uc.FileSetCursor((uc.file.cursor+len(uc.file.trajectory)-1)%len(uc.file.trajectory), c)
	if err != nil {
		return
	}
	return uc.FileCurrent(c)
}

func (uc *UseCase) FileSetCursor(i int, c chan string) (err error) {
	if uc.file.trajectory == nil {
		return errors.New("trajectory is empty")
	}
	if len(uc.file.trajectory) < i {
		return errors.New("invalid cursor value")
	}
	uc.file.cursor = i
	return
}

func (uc *UseCase) FileRun(c chan string) (err error) {
	for true {
		err = uc.FileNext(c)
		if err != nil {
			return
		}
		if uc.file.cursor == 0 {
			break
		}
	}
	return
}
