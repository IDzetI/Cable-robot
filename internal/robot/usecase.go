package robot

import (
	"errors"
	"github.com/IDzetI/Cable-robot/internal/robot/controller"
	robot_extruder "github.com/IDzetI/Cable-robot/internal/robot/controller/extruder"
	"github.com/IDzetI/Cable-robot/internal/robot/kinematics"
	"github.com/IDzetI/Cable-robot/internal/robot/parser"
	"github.com/IDzetI/Cable-robot/internal/robot/trajectory"
)

type UseCase struct {
	controller               robot_controller.Controller
	trajectoryCartesianSpace robot_trajectory.Trajectory
	trajectoryJoinSpace      robot_trajectory.Trajectory
	kinematics               robot_kinematics.Kinematics
	extruder                 robot_extruder.Extruder
	file                     file
}

func New(c robot_controller.Controller,
	tc, tj robot_trajectory.Trajectory,
	k robot_kinematics.Kinematics,
	e robot_extruder.Extruder) (uc UseCase) {
	return UseCase{
		controller:               c,
		trajectoryCartesianSpace: tc,
		trajectoryJoinSpace:      tj,
		kinematics:               k,
		extruder:                 e,
		file: file{
			trajectory: [][]float64{},
			cursor:     0,
			plt:        nil,
			tr:         robot_parser.Rt{},
		},
	}
}

func (uc *UseCase) ConfigPLT(up, down float64, start []float64) (err error) {
	if len(start) != 3 {
		return errors.New("invalid start point")
	}
	uc.file.plt = &robot_parser.Plt{
		Up:    up,
		Down:  down,
		Start: start,
	}
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
