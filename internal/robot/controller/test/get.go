package robot_controller_test

import "log"

func (c *controller) GetDegrees() (degrees []float64, err error) {
	log.Println("controller get degree")
	degrees = []float64{100, 100, 100, 103}
	return
}

func (c *controller) GetPeriod() (period float64) {
	return 4e-3
}
