package robot

import (
	"github.com/IDzetI/Cable-robot/internal/robot/controller"
	"github.com/IDzetI/Cable-robot/internal/robot/extruder"
	"github.com/IDzetI/Cable-robot/internal/robot/kinematics"
	"github.com/IDzetI/Cable-robot/internal/robot/trajectory"
	"sync"
)

type UseCase struct {
	controller               robot_controller.Controller
	trajectoryCartesianSpace robot_trajectory.Trajectory
	trajectoryJoinSpace      robot_trajectory.Trajectory
	kinematics               robot_kinematics.Kinematics
	extruder                 robot_extruder.Extruder
	file                     *file

	shift       []float64
	position    []float64
	movingMutex sync.Mutex
	resetFlag   bool
	stopFlag    bool

	commands chan func()
}

func New(c robot_controller.Controller,
	tc, tj robot_trajectory.Trajectory,
	k robot_kinematics.Kinematics,
	e robot_extruder.Extruder) (uc UseCase) {
	uc = UseCase{
		controller:               c,
		trajectoryCartesianSpace: tc,
		trajectoryJoinSpace:      tj,
		kinematics:               k,
		extruder:                 e,
		commands:                 make(chan func(), 1024),
	}
	go uc.execute()
	return
}

func (uc *UseCase) ReadDegrees() (degrees []float64, err error) {
	return uc.controller.GetDegrees()
}

func (uc *UseCase) ControlOn(c chan string) (err error) {
	return uc.controller.ControlON()
}

func (uc *UseCase) ControlOff(c chan string) (err error) {
	return uc.controller.ControlOFF()
}
