package robot

import (
	"github.com/IDzetI/Cable-robot/internal/robot/controller"
	"github.com/IDzetI/Cable-robot/internal/robot/kinematics"
	"github.com/IDzetI/Cable-robot/internal/robot/trajectory"
)

type UseCase struct {
	controller               robot_controller.Controller
	trajectoryCartesianSpace robot_trajectory.Trajectory
	trajectoryJoinSpace      robot_trajectory.Trajectory
	kinematics               robot_kinematics.Kinematics
	file                     file
}

func (uc *UseCase) ReadDegrees() (lengths []float64, err error) {
	return uc.controller.GetDegrees()
}

func (uc *UseCase) ControlOn(c chan string) (err error) {
	return uc.controller.ControlON()
}

func (uc *UseCase) ControlOff(c chan string) (err error) {
	return uc.controller.ControlOFF()
}
