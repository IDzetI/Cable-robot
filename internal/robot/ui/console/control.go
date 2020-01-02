package robot_ui_console

import (
	"bufio"
	"fmt"
	"github.com/IDzetI/Cable-robot/internal/robot"
	"log"
	"os"
	"strconv"
	"strings"
)

func Start(uc robot.UseCase) (err error) {

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
			fmt.Println(uc.ReadDegree())

		case cmdSet:
			switch data[1] {
			case cmdSetSpeed:
				setSingleValue(data[2], uc.SetSpeed)
				break
			case cmdSetMinSpeed:
				setSingleValue(data[2], uc.SetMinSpeed)
				break
			case cmdSetAcceleration:
				setSingleValue(data[2], uc.SetAcceleration)
				break
			case cmdSetDeceleration:
				setSingleValue(data[2], uc.SetDeceleration)
				break
			default:
				break
			}
			break
		case cmdInit:
			moveToPoint(data[1:], uc.MoveInJoinSpace)
			break
		case cmdMove:
			moveToPoint(data[1:], uc.MoveInCartesianSpace)
			break
		case cmdControl:
			switch data[1] {
			case cmdControlON:
				exec(uc.ControlOn)
				break
			case cmdControlOFF:
				exec(uc.ControlOff)
				break
			case cmdControlRESET:
				exec(uc.ControlOff)
				exec(uc.ControlOn)
			}
			break
		case cmdFile:
			switch data[1] {
			case cmdFilePase:
				execWithString(data[2], uc.FileParse)
				break
			case cmdFileLoad:
				execWithString(data[2], uc.FileLoad)
				break
			case cmdFileInit:
				exec(uc.FileInit)
				break
			case cmdFileNext:
				exec(uc.FileNext)
				break
			case cmdFileCurrent:
				exec(uc.FileCurrent)
				break
			case cmdFilePrevious:
				exec(uc.FilePrevious)
				break
			case cmdFileSetCursor:
				execWithInt(data[2], uc.FileSetCursor)
				break
			case cmdFileGo:
				exec(uc.FileGo)
				break
			case cmdFileRun:
				exec(uc.FileRun)
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

func exec(f func() error) {
	if err := f(); err != nil {
		log.Println(err.Error())
	}
}

func execWithString(s string, f func(string) error) {
	if err := f(s); err != nil {
		log.Println(err.Error())
	}
}

func execWithInt(s string, f func(int) error) {
	if v, err := strconv.Atoi(s); err == nil {
		if err := f(v); err != nil {
			log.Println(err.Error())
		}
	} else {
		log.Println(err.Error())
	}
}

func moveToPoint(strPoint []string, f func([]float64) error) {
	if len(strPoint) != 3 {
		log.Println("invalid position")
		return
	}
	var point []float64
	for _, s := range strPoint {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			log.Println(err.Error())
			return
		}
		point = append(point, v)
	}
	err := f(point)
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
