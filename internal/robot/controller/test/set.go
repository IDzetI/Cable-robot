package robot_controller_test

import (
	"log"
)

func (c *controller) SendTrajectory(lengths [][]float64) (err error) {
	panic("implement me")
	//TODO
}

func (c *controller) SetDegrees(degrees []float64) (err error) {
	log.Println("controller set degree", degrees)
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
