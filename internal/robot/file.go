package robot

import (
	"errors"
	"github.com/IDzetI/Cable-robot/internal/robot/parser"
	"strconv"
	"strings"
)

type file struct {
	trajectory [][]float64
	cursor     int
	plt        *robot_parser.Plt
	tr         robot_parser.Rt
}

func (u *UseCase) ConfigPLT(up, down float64, start []float64) (err error) {
	if len(start) != 3 {
		return errors.New("invalid start point")
	}
	u.file.plt = &robot_parser.Plt{
		Up:    up,
		Down:  down,
		Start: start,
	}
	return
}

func (u *UseCase) FileLoad(fileString string, c chan string) (err error) {

	u.file = &file{}

	fileName := strings.Split(fileString, ".")

	var trajectory [][]float64

	switch fileName[len(fileName)-1] {
	case "plt":
		if u.file.plt == nil {
			err = errors.New("plt config is empty")
			return
		}
		trajectory, err = u.file.plt.Read(fileString, c)

	case "rt":
		trajectory, err = u.file.tr.Read(fileString, c)

	default:
		err = errors.New("incorrect file extension")
		return
	}

	if err != nil {
		return
	}
	u.file.trajectory = trajectory
	u.file.cursor = 0
	return
}

func (u *UseCase) FileInit(c chan string) (err error) {
	return u.MoveInJoinSpace(u.file.trajectory[u.file.cursor], c)
}

func (u *UseCase) FileNext(c chan string) (err error) {
	err = u.FileSetCursor((u.file.cursor+1)%len(u.file.trajectory), c)
	if err != nil {
		return
	}
	return u.FileCurrent(c)
}

func (u *UseCase) FileCurrent(c chan string) (err error) {
	return u.MoveInCartesianSpace(u.file.trajectory[u.file.cursor], c)
}

func (u *UseCase) FilePrevious(c chan string) (err error) {
	err = u.FileSetCursor((u.file.cursor+len(u.file.trajectory)-1)%len(u.file.trajectory), c)
	if err != nil {
		return
	}
	return u.FileCurrent(c)
}

func (u *UseCase) FileSetCursor(i int, c chan string) (err error) {
	if u.file.trajectory == nil {
		return errors.New("trajectory is empty")
	}
	if len(u.file.trajectory) < i {
		return errors.New("invalid cursor value")
	}
	u.file.cursor = i
	c <- "Cursor = " + strconv.Itoa(u.file.cursor)
	return
}

func (u *UseCase) FileRun(c chan string) (err error) {
	for true {
		err = u.FileNext(c)
		if err != nil {
			return
		}
		if u.file.cursor == 0 {
			break
		}
	}
	return
}
