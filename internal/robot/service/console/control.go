package robot_service_console

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

	for exit := false; !exit; {
		//read line
		line, err := reader.ReadString('\n')
		if err != nil {
			break
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

			case cmdSetMinSpeed:
				setSingleValue(data[2], uc.SetMinSpeedCartesianSpace)

			case cmdSetAcceleration:
				setSingleValue(data[2], uc.SetAccelerationCartesianSpace)

			case cmdSetDeceleration:
				setSingleValue(data[2], uc.SetDecelerationCartesianSpace)

			default:
			}

		case cmdInit:
			moveToPoint(data[1:], c, uc.MoveInJoinSpace)

		case cmdMove:
			moveToPoint(data[1:], c, uc.MoveInCartesianSpace)

		case cmdControl:
			switch data[1] {

			case cmdControlON:
				exec(c, uc.ControlOn)

			case cmdControlOFF:
				exec(c, uc.ControlOff)

			case cmdControlRESET:
				exec(c, uc.ControlOff)
				exec(c, uc.ControlOn)
			}

		case cmdFile:
			switch data[1] {

			case cmdFileLoad:
				execWithString(data[2], c, uc.FileLoad)

			case cmdFileInit:
				exec(c, uc.FileInit)

			case cmdFileNext:
				exec(c, uc.FileNext)

			case cmdFileCurrent:
				exec(c, uc.FileCurrent)

			case cmdFilePrevious:
				exec(c, uc.FilePrevious)

			case cmdFileSetCursor:
				execWithInt(data[2], c, uc.FileSetCursor)

			case cmdFileGo:
				fileGo(reader, uc, c)

			case cmdFileRun:
				exec(c, uc.FileRun)

			}

		case cmdExit:
			exit = true
			break

		default:

		}
	}
	return
}

func fileGo(reader *bufio.Reader, uc robot.UseCase, c chan string) {
	fmt.Println("press ENTER to next\n" +
		"write p and press ENTER to go to previous position\n" +
		"write q and press ENTER to exit from this mode")
	for true {
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err.Error())
			return
		}
		switch command {
		case cmdFileGoPrevious:
			exec(c, uc.FilePrevious)
		case cmdFileGoExit:
			return
		default:
			exec(c, uc.FileNext)
		}
	}
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
