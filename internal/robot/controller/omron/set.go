package robot_controller_omron

import (
	"fmt"
	"github.com/IDzetI/Cable-robot/pkg/utils"
	"time"
)

func (c *controller) SendTrajectory(lengths [][]float64) (err error) {
	steps := len(lengths)

	err = c.SetDegrees(lengths[0])
	if err != nil {
		return
	}

	err = c.ControlON()
	if err != nil {
		return
	}

	//create timer
	period := time.Duration(c.period) * time.Second
	done := make(chan bool, 1)
	ticker := time.NewTicker(period)

	// current line
	counter := 0
	fmt.Println("robot move")
	go func() {
		for counter < steps {
			select {
			case <-ticker.C:
				err = c.SetDegrees(lengths[counter])
				if err != nil {
					done <- false
					return
				}
				counter++
			}
		}
		done <- true
	}()
	<-done
	return
}

func (c *controller) SetDegrees(degrees []float64) (err error) {

	//parse lengths
	var data []uint16
	for _, length := range degrees {
		for _, d := range utils.Float64Uint16(length) {
			data = append(data, d)
		}
	}

	//write data
	return c.robot.WriteDNoResponse(cEncodersAddress, data)
}

func (c *controller) ControlON() (err error) {
	return c.robot.WriteD(controlAddress, []uint16{1})
}

func (c *controller) ControlOFF() (err error) {
	return c.robot.WriteD(controlAddress, []uint16{0})
}

func (c *controller) HasError() (e bool, err error) {
	isError, err := c.robot.ReadD(errorAddress, 1)
	if err != nil {
		return
	}

	e = isError[0] == 1
	return
}
