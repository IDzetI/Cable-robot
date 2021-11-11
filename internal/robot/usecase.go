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
	speed       []float64
	movingMutex sync.Mutex
	resetFlag   bool
	stopFlag    bool

	commands             chan func()
	externalInterception int64
}

func New(c robot_controller.Controller,
	tc, tj robot_trajectory.Trajectory,
	k robot_kinematics.Kinematics) (uc UseCase) {
	uc = UseCase{
		controller:               c,
		trajectoryCartesianSpace: tc,
		trajectoryJoinSpace:      tj,
		kinematics:               k,
		commands:                 make(chan func(), 1024),
	}
	go uc.execute()
	return
}

func (u *UseCase) GetPosition() []float64 {
	if u.position == nil {
		return []float64{0, 0, 0}
	} else {
		return u.position
	}
}

func (u *UseCase) GetSpeed() []float64 {
	if u.speed == nil {
		return []float64{0, 0, 0}
	} else {
		return u.speed
	}
}

func (u *UseCase) ReadDegrees() (degrees []float64, err error) {
	return u.controller.GetDegrees()
}

func (u *UseCase) ControlOn(c chan string) (err error) {
	c <- "Control starting"
	return u.controller.ControlON()
}

func (u *UseCase) ControlOff(c chan string) (err error) {
	c <- "Control off"
	return u.controller.ControlOFF()
}
