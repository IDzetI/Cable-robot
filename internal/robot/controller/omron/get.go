package robot_controller_omron

import (
	"github.com/IDzetI/Cable-robot.git/pkg/utils"
)

func (c *controller) GetLengths() (lengths []float64, err error) {

	bytes, err := c.robot.ReadD(vEncodersAddress, cableNumber*floatLength)
	if err != nil {
		return
	}

	for i := 0; i < cableNumber; i++ {
		lengths = append(lengths, utils.Uint16Float64(bytes[i*floatLength:(i+1)*floatLength]))
	}
	return
}
