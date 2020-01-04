package robot_ui_console

import (
	"bufio"
	"fmt"
	"github.com/IDzetI/Cable-robot/internal/robot"
	"github.com/IDzetI/Cable-robot/pkg/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

func Start(uc robot.UseCase) (err error) {

	//create post back listener
	c := make(chan string)
	go listener(c)

	//create reader for reed from console
	reader := bufio.NewReader(os.Stdin)

	for true {
		//read line
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		line = strings.ReplaceAll(strings.ReplaceAll(line, "\n", ""), "\r", "")
		data := strings.Split(line, " ")

		switch data[0] {
		case cmdReadDegree:
			lengths, err := uc.ReadDegrees()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(utils.ToString(lengths))
			}

		case cmdSet:
			switch data[1] {
			case cmdSetSpeed:
				setSingleValue(data[2], uc.SetSpeedCartesianSpace)
				break
			case cmdSetMinSpeed:
				setSingleValue(data[2], uc.SetMinSpeedCartesianSpace)
				break
			case cmdSetAcceleration:
				setSingleValue(data[2], uc.SetAccelerationCartesianSpace)
				break
			case cmdSetDeceleration:
				setSingleValue(data[2], uc.SetDecelerationCartesianSpace)
				break
			default:
				break
			}
			break
		case cmdInit:
			moveToPoint(data[1:], c, uc.MoveInJoinSpace)
			break
		case cmdMove:
			moveToPoint(data[1:], c, uc.MoveInCartesianSpace)
			break
		case cmdControl:
			switch data[1] {
			case cmdControlON:
				exec(c, uc.ControlOn)
				break
			case cmdControlOFF:
				exec(c, uc.ControlOff)
				break
			case cmdControlRESET:
				exec(c, uc.ControlOff)
				exec(c, uc.ControlOn)
			}
			break
		case cmdFile:
			switch data[1] {
			case cmdFilePase:
				execWithString(data[2], c, uc.FileLoadPLT)
				break
			case cmdFileLoad:
				execWithString(data[2], c, uc.FileLoad)
				break
			case cmdFileInit:
				exec(c, uc.FileInit)
				break
			case cmdFileNext:
				exec(c, uc.FileNext)
				break
			case cmdFileCurrent:
				exec(c, uc.FileCurrent)
				break
			case cmdFilePrevious:
				exec(c, uc.FilePrevious)
				break
			case cmdFileSetCursor:
				execWithInt(data[2], c, uc.FileSetCursor)
				break
			case cmdFileGo:
				exec(c, uc.FileGo)
				break
			case cmdFileRun:
				exec(c, uc.FileRun)
				break
			}
			break
		case cmdExit:
			return
		default:
			break
		}
	}
	return
}

func exec(c chan string, f func(chan string) error) {
	if err := f(c); err != nil {
		log.Println(err.Error())
	}
}

func execWithString(s string, c chan string, f func(string, chan string) error) {
	if err := f(s, c); err != nil {
		log.Println(err.Error())
	}
}

func execWithInt(s string, c chan string, f func(int, chan string) error) {
	if v, err := strconv.Atoi(s); err == nil {
		if err := f(v, c); err != nil {
			log.Println(err.Error())
		}
	} else {
		log.Println(err.Error())
	}
}

func moveToPoint(strPoint []string, c chan string, f func([]float64, chan string) error) {
	var point []float64
	for _, s := range strPoint {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Println(err.Error())
			return
		}
		point = append(point, v)
	}
	err := f(point, c)
	if err != nil {
		log.Println(err)
		return
	}
}

func setSingleValue(s string, f func(float64) error) {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = f(value)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func listener(c chan string) {
	for msg := range c {
		log.Println(msg)
	}
}
