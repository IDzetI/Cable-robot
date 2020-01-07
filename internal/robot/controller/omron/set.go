package robot_controller_omron

import (
	"github.com/IDzetI/Cable-robot.git/pkg/utils"
)

func (c *controller) SetLengths(lengths []float64) (err error) {

	//parse lengths
	var data []uint16
	for _, length := range lengths {
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
