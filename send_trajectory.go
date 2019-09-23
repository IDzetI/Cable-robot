package main

import (
	"cable_robot_v2/fins"
	"fmt"
	"time"
)

/*
Function send one robot package per period
*/
func sendTrajectory(robot *fins.Client, packages [][]uint16) {
	rows := len(packages)

	sendPackage(robot,packages[0])
	ONControlMode(robot)

	//create timer
	period := 4*time.Millisecond
	done := make(chan bool, 1)
	ticker := time.NewTicker(period)

	// current line
	counter := 0
	fmt.Println("robot move")
	go func () {
		for {
			select {
			case <- ticker.C:
				//TODO stop from console
				if counter >= rows {
					done <- true
					break
				}
				//send one package
				sendPackage(robot,packages[counter])
				counter ++
			}
		}
	}()
	<- done
	fmt.Println("robot stop")
}
