package robot_controller_test

import (
	"log"
	"time"
)

func (c *controller) SendTrajectory(degrees [][]float64) (err error) {
	log.Println("controller set degrees", degrees)
	return
}

func (c *controller) SetDegrees(degrees []float64) (err error) {
	log.Println(time.Now(), "controller set degree", degrees)
	return
}

func (c *controller) ControlON() (err error) {
	log.Println("controller control on")
	return
}

func (c *controller) ControlOFF() (err error) {
	log.Println("controller control off")
	return
}

func (c *controller) HasError() (e bool, err error) {
	log.Println("controller has error")
	return
}

func (c *controller) Reset() (err error) {
	log.Println("controller was reset")
	return
}
