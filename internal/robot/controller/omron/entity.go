package robot_controller_omron

import (
	"github.com/IDzetI/Cable-robot/pkg/fins"
)

type controller struct {
	robot  *fins.Client
	period float64
}

func New(period float64, address string) (controller controller, err error) {
	controller.robot, err = fins.NewClient(address)
	if err != nil {
		return
	}
	controller.period = period
	return
}
