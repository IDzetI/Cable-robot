package robot_controller_omron

import (
	robot_controller "github.com/IDzetI/Cable-robot/internal/robot/controller"
	"github.com/IDzetI/Cable-robot/pkg/fins"
)

type controller struct {
	robot  *fins.Client
	period float64
}

func New(period float64, address string) (c robot_controller.Controller, err error) {
	contr := controller{
		period: period,
	}
	contr.robot, err = fins.NewClient(address)
	if err != nil {
		return
	}
	return &contr, nil
}
