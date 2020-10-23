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
		case cmdStop:
			go uc.Stop(c)

		case cmdResume:
			go uc.Resume(c)

		case cmdReset:
			go uc.Reset(c)

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
				go setSingleValue(data[2], uc.SetSpeedCartesianSpace)

			case cmdSetMinSpeed:
				go setSingleValue(data[2], uc.SetMinSpeedCartesianSpace)

			case cmdSetAcceleration:
				go setSingleValue(data[2], uc.SetAccelerationCartesianSpace)

			case cmdSetDeceleration:
				go setSingleValue(data[2], uc.SetDecelerationCartesianSpace)

			case cmdSetExtruderSpeed:
				go setSingleValue(data[2], uc.SetExtruderSpeed)

			case cmdSetShift:
				go setArrayValue(data[2:], uc.SetShift)
			}

		case cmdSetJoinSpace:
			switch data[1] {

			case cmdSetSpeed:
				go setSingleValue(data[2], uc.SetSpeedJoinSpace)

			case cmdSetMinSpeed:
				go setSingleValue(data[2], uc.SetMinSpeedJoinSpace)

			case cmdSetAcceleration:
				go setSingleValue(data[2], uc.SetAccelerationJoinSpace)

			case cmdSetDeceleration:
				go setSingleValue(data[2], uc.SetDecelerationJoinSpace)
			}

		case cmdInit:
			go moveToPoint(data[1:], c, uc.MoveInJoinSpace)

		case cmdMove:
			go moveToPoint(data[1:], c, uc.MoveInCartesianSpace)

		case cmdControl:
			switch data[1] {

			case cmdControlON:
				go exec(c, uc.ControlOn)

			case cmdControlOFF:
				go exec(c, uc.ControlOff)

			case cmdControlRESET:
				go exec(c, uc.ControlOff)
				go exec(c, uc.ControlOn)
			}

		case cmdFile:
			switch data[1] {

			case cmdFileLoad:
				go execWithString(data[2], c, uc.FileLoad)

			case cmdFileInit:
				go exec(c, uc.FileInit)

			case cmdFileNext:
				go exec(c, uc.FileNext)

			case cmdFileCurrent:
				go exec(c, uc.FileCurrent)

			case cmdFilePrevious:
				go exec(c, uc.FilePrevious)

			case cmdFileSetCursor:
				go execWithInt(data[2], c, uc.FileSetCursor)

			case cmdFileGo:
				fileGo(reader, uc, c)

			case cmdFileRun:
				go exec(c, uc.FileRun)

			}

		case cmdExit:
			exit = true

		default:

		}
	}
	return
}

func fileGo(reader *bufio.Reader, uc robot.UseCase, c chan string) {
	fmt.Println("press ENTER to next\n" +
		"write p and press ENTER to go to previous position\n" +
		"write q and press ENTER to exit from this mode")
	for {
		command, err := reader.ReadString('\n')
		c <- command
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
	point, err := utils.ToFloatArray(strPoint)
	if err != nil {
		log.Println(err)
		return
	}
	err = f(point, c)
	if err != nil {
		log.Println(err)
		return
	}
}

func setArrayValue(strValues []string, f func([]float64) error) {
	values, err := utils.ToFloatArray(strValues)
	if err != nil {
		log.Println(err)
		return
	}
	err = f(values)
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
	} else {
		log.Println("OK")
	}
}

func listener(c chan string) {
	for msg := range c {
		log.Println(msg)
	}
}
