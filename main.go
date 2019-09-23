package main

import (
	"cable_robot_v2/fins"
)


func main() {
	plcAddr := "192.168.250.1:9600"

	robot := fins.NewClient(plcAddr)
	defer robot.CloseConnection()

	//degrees := readCurrentDegrees(robot)
	//for i,_ := range degrees{
	//	degrees[i] += 10
	//}
	//ONControlMode(robot)
	//sendPackage(robot,PacToUint([4]float64{-100.0,-100.0,-100.0,-100.0}))
	//sendPackage(robot,PacToUint(degrees))
	//OFFControlMode(robot)
	consolePositionControlWithSmoothMoving(robot)
}
