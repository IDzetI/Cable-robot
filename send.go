package main

import (
	"github.com/IDzetI/Cable-robot/pkg/fins"
	"log"
	"strings"
)

/*
Function read set on engels from file and send they to robot
*/
func sendDegreesFromFile(robot *fins.Client, fileName string) {
	fileLines := strings.Split(ReadAll(fileName), "\r\n")
	for i := 0; i < len(fileLines); i++ {
		setAngles(robot, LineToAngles(fileLines[i]))
	}
}

/*
Functions send package to robot
*/
func sendPackage(robot *fins.Client, pac []uint16) {
	err := robot.WriteDNoResponse(420, pac)
	if err != nil {
		log.Fatal(err)
	}
}

/*
Function send 4 angles to robot
*/
func setAngles(robot *fins.Client, angles [4]float64) {
	sendPackage(robot, PacToUint(angles))
}

/*
Function convert angels package to robot package
*/
func PacToUint(pac [4]float64) []uint16 {
	res := make([]uint16, 4*len(pac))

	for i := 0; i < len(pac); i++ {
		current := Float64Uint16(pac[i])
		for j := 0; j < len(current); j++ {
			res[4*i+j] = current[j]
		}
	}
	return res
}
