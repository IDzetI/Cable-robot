package robot_service_console

import (
	"bufio"
	"fmt"
	"github.com/IDzetI/Cable-robot/internal/robot"
	robot_service "github.com/IDzetI/Cable-robot/internal/robot/service"
	"github.com/IDzetI/Cable-robot/pkg/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

type service struct {
	uc *robot.UseCase
}

func New(uc *robot.UseCase) robot_service.Service {
	s := service{uc: uc}
	return &s
}

func (s *service) Start() (err error) {

	//create post back listener
	c := make(chan string)
	go s.listener(c)

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
		case "e":
			err = s.uc.ExternalControl([]float64{0, 500, 0})
			if err != nil {
				panic(err)
			}

		case cmdStop:
			go s.uc.Stop(c)

		case cmdResume:
			go s.uc.Resume(c)

		case cmdReset:
			go s.uc.Reset(c)

		case cmdReadDegree:
			lengths, err := s.uc.ReadDegrees()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(utils.ToString(lengths))
			}

		case cmdSet:
			switch data[1] {

			case cmdSetSpeed:
				go setSingleValue(data[2], s.uc.SetSpeedCartesianSpace)

			case cmdSetMinSpeed:
				go setSingleValue(data[2], s.uc.SetMinSpeedCartesianSpace)

			case cmdSetAcceleration:
				go setSingleValue(data[2], s.uc.SetAccelerationCartesianSpace)

			case cmdSetDeceleration:
				go setSingleValue(data[2], s.uc.SetDecelerationCartesianSpace)

			case cmdSetExtruderSpeed:
				go setSingleValue(data[2], s.uc.SetExtruderSpeed)

			case cmdSetShift:
				go setArrayValue(data[2:], s.uc.SetShift)
			}

		case cmdSetJoinSpace:
			switch data[1] {

			case cmdSetSpeed:
				go setSingleValue(data[2], s.uc.SetSpeedJoinSpace)

			case cmdSetMinSpeed:
				go setSingleValue(data[2], s.uc.SetMinSpeedJoinSpace)

			case cmdSetAcceleration:
				go setSingleValue(data[2], s.uc.SetAccelerationJoinSpace)

			case cmdSetDeceleration:
				go setSingleValue(data[2], s.uc.SetDecelerationJoinSpace)
			}

		case cmdInit:
			go moveToPoint(data[1:], c, s.uc.MoveInJoinSpace)

		case cmdMove:
			go moveToPoint(data[1:], c, s.uc.MoveInCartesianSpace)

		case cmdControl:
			switch data[1] {

			case cmdControlON:
				go exec(c, s.uc.ControlOn)

			case cmdControlOFF:
				go exec(c, s.uc.ControlOff)

			case cmdControlRESET:
				go exec(c, s.uc.ControlOff)
				go exec(c, s.uc.ControlOn)
			}

		case cmdFile:
			switch data[1] {

			case cmdFileLoad:
				go execWithString(data[2], c, s.uc.FileLoad)

			case cmdFileInit:
				go exec(c, s.uc.FileInit)

			case cmdFileNext:
				go exec(c, s.uc.FileNext)

			case cmdFileCurrent:
				go exec(c, s.uc.FileCurrent)

			case cmdFilePrevious:
				go exec(c, s.uc.FilePrevious)

			case cmdFileSetCursor:
				go execWithInt(data[2], c, s.uc.FileSetCursor)

			case cmdFileGo:
				s.fileGo(reader, c)

			case cmdFileRun:
				go exec(c, s.uc.FileRun)

			}

		case cmdExit:
			exit = true

		default:

		}
	}
	return
}

func (s *service) fileGo(reader *bufio.Reader, c chan string) {
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
			exec(c, s.uc.FilePrevious)
		case cmdFileGoExit:
			return
		default:
			exec(c, s.uc.FileNext)
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

func (s *service) listener(c chan string) {
	for msg := range c {
		log.Println(msg)
	}
}
