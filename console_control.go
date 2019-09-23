package main

import (
	"bufio"
	"cable_robot_v2/fins"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Function to control position of endeffektor with speed control
*/
func consolePositionControlWithSmoothMoving(robot *fins.Client) {

	//set default values
	speed := 300.0
	acceleration := 100.0
	deceleration := 100.0
	minSpeed := 5.0
	position := [3]float64{0, 0, 0}
	var cordFile [][3]float64
	cordFilePointer := 0

	//create reader for reed from console
	reader := bufio.NewReader(os.Stdin)

	//main loop
	for true {
		//read line
		line, err := reader.ReadString('\n')
		check(err)

		line = strings.ReplaceAll(strings.ReplaceAll(line, "\n", ""), "\r", "")
		data := strings.Split(line, " ")

		switch data[0] {
		case "dnow":
			readCurrentDegrees(robot)
		case "set":
			switch data[1] {
			case "cur":
				position = getPosition(data[2:5], position, "current position = ")
				break
			case "speed":
				speed = getValue(data[2], speed, "speed", 1, 500)
				break
			case "acc":
				acceleration = getValue(data[2], speed, "acceleration", 1, 1000)
				break
			case "dec":
				deceleration = getValue(data[2], deceleration, "deceleration", 1, 1000)
				break
			case "minspeed":
				minSpeed = getValue(data[2], minSpeed, "minimum speed", 0, 10)
				break
			default:
				break
			}
			break
		case "init":
			position = movementInJointSpace(robot, getPosition(data[1:4], position, "move in joint space to position = "), speed, acceleration, deceleration, minSpeed)
		case "move":
			position = movementInCartesianSpace(robot, position, getPosition(data[1:4], position, "move in cartesian space to position = "), speed, acceleration, deceleration, minSpeed)
			break
		case "stop":
			stopMotors(robot)
			break
		case "start":
			turnONMotors(robot)
			break
		case "control":
			switch data[1] {
			case "on":
				ONControlMode(robot)
				break
			case "off":
				OFFControlMode(robot)
				break
			case "reset":
				OFFControlMode(robot)
				ONControlMode(robot)
			}
			break
		case "f":
			switch data[1] {
			case "parse":
				loadPltFile(/*"trajectories/plt/"+*/data[2]+".plt", /*"testing/target_data/"+*/data[2]+".txt", [3]float64{0, 0, 0})
				break
			case "load":
				cordFile = loadCordFile(/*"testing/target_data/" + */data[2] + ".txt")
				break
			case "init":
				fmt.Println("moving to ", cordFile[cordFilePointer])
				position = movementInJointSpace(robot, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
				break
			case "next":
				cordFilePointer = (cordFilePointer + 1) % len(cordFile)
				fmt.Println("moving to ", cordFile[cordFilePointer])
				position = movementInCartesianSpace(robot, position, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
				break
			case "cur":
				fmt.Println("moving to ", cordFile[cordFilePointer])
				position = movementInCartesianSpace(robot, position, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
				break
			case "prev":
				cordFilePointer = (cordFilePointer + (len(cordFile) - 1)) % len(cordFile)
				fmt.Println("moving to ", cordFile[cordFilePointer])
				position = movementInCartesianSpace(robot, position, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
				break
			case "set":
				cordFilePointer = int(getValue(data[2], float64(cordFilePointer), "current point in coordinate file", 0, math.MaxInt64))
				break
			case "go":
				for i := cordFilePointer; i < len(cordFile); i++ {
					cordFilePointer = (cordFilePointer + 1) % len(cordFile)
					fmt.Println("moving to ", cordFilePointer, cordFile[cordFilePointer])
					position = movementInCartesianSpace(robot, position, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
					readCurrentDegrees(robot)
					waitEnter()
				}
				break
			case "run":
				for i := cordFilePointer; i < len(cordFile)-1; i++ {
					cordFilePointer = (cordFilePointer + 1) % len(cordFile)
					fmt.Println("moving to ", cordFilePointer, cordFile[cordFilePointer])
					position = movementInCartesianSpace(robot, position, cordFile[cordFilePointer], speed, acceleration, deceleration, minSpeed)
					readCurrentDegrees(robot)
					if readError(robot) {
						waitEnter()
					}
				}
				break
			}
			break
		case "exit":
			return
		default:
			break
		}
	}
}

/*
Function to to send angles from console to robot.
*/
func consoleAnglesControl(robot *fins.Client) {
	//create reader for read from console
	reader := bufio.NewReader(os.Stdin)

	for true {
		//read line
		fmt.Println("Enter robot angles  (in degree like <11.1 22.2 33.3 44.4>) or q for exit:")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		//exit condition input = 'q'
		if strings.Compare(text, "q") == 0 {
			break
		}

		//sent angles
		setAngles(robot, LineToAngles(text))

		//new line
		fmt.Println()
	}
}

/*
Function to control position of endeffektor
*/
func consoleCoordinateControl(robot *fins.Client) {
	//create reader for reed from console
	reader := bufio.NewReader(os.Stdin)

	//infinite loop
	for true {
		//read line
		fmt.Println("Enter robot x y z (in mm like <11.1 22.2 33.3>) or q for exit:")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		//exit condition input = 'q'
		if strings.Compare(text, "q") == 0 {
			break
		}

		//check number of inputs
		values := strings.Split(text, " ")
		if len(values) != 3 {
			fmt.Println("ERROR Please enter correct x y z")
			return
		}

		x := [3]float64{}
		for i := 0; i < 3; i++ {
			//convert strings to float64
			buff, err := strconv.ParseFloat(values[i], 64)
			if err != nil {
				fmt.Println(err)
			}
			x[i] = buff
		}

		//set coordinates
		setAngles(robot, xyzToDegrees(x))
	}
}
