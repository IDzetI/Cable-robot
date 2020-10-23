package robot_controller_omron

import (
	"github.com/IDzetI/Cable-robot/pkg/utils"
)

func (c *controller) GetDegrees() (degrees []float64, err error) {

	bytes, err := c.robot.ReadD(vEncodersAddress, cableNumber*floatLength)
	if err != nil {
		return
	}

	for i := 0; i < cableNumber; i++ {
		degrees = append(degrees, utils.Uint16Float64(bytes[i*floatLength:(i+1)*floatLength]))
	}
	return
}

func (c *controller) GetPeriod() (period float64) {
	return c.period
}
