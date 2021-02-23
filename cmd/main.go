package main

import (
	"github.com/IDzetI/Cable-robot/config"
	"github.com/IDzetI/Cable-robot/internal/robot"
	robot_controller_test "github.com/IDzetI/Cable-robot/internal/robot/controller/test"
	"github.com/IDzetI/Cable-robot/internal/robot/kinematics/rw_model"
	robot_service_rest "github.com/IDzetI/Cable-robot/internal/robot/service/rest"
	"github.com/IDzetI/Cable-robot/internal/robot/trajectory/v1"
	"log"
)

func main() {
	log.Println("Program starting...")

	//read config from config.yaml
	conf, err := config.Init()
	if err != nil {
		panic(err)
	}

	//initialise robot controller
	//robotController, err := robot_controller_omron.New(conf.Period,conf.Ip)
	robotController, err := robot_controller_test.New()
	if err != nil {
		panic(err)
	}

	//initialise robot kinematic model
	var robotMotors []robot_kinematics_rw_model.Motor
	for _, motor := range conf.Motors {
		robotMotors = append(robotMotors, robot_kinematics_rw_model.Motor{
			DrumH:     motor.Drum.H,
			DrumR:     motor.Drum.R,
			RollerR:   motor.RollerRadius,
			ExitPoint: motor.ExitPoint,
		})
	}
	robotKinematics, err := robot_kinematics_rw_model.New(robotMotors)
	if err != nil {
		panic(err)
	}

	//initialise trajectory planing in cartesian space
	robotCartesianSpaceTrajectory, err := robot_trajectory_v1.New(
		conf.CartesianSpace.Speed,
		conf.CartesianSpace.MinSpeed,
		conf.CartesianSpace.Acceleration,
		conf.CartesianSpace.Deceleration,
		conf.Period,
		conf.Workspace,
	)
	if err != nil {
		panic(err)
	}

	//initialise trajectory
	robotJoinSpaceTrajectory, err := robot_trajectory_v1.New(
		conf.JoinSpace.Speed,
		conf.JoinSpace.MinSpeed,
		conf.JoinSpace.Acceleration,
		conf.JoinSpace.Deceleration,
		conf.Period,
		[][]float64{
			{-5000, 5000},
			{-5000, 5000},
			{-5000, 5000},
			{-5000, 5000},
		},
	)
	if err != nil {
		panic(err)
	}

	//initialise extruder
	//robotExtruder, err := robot_extruder_smsd.New(conf.Extruder.Port)
	//if err != nil {
	//	panic(err)
	//}

	//initialise robot
	robotUseCase := robot.New(
		robotController,
		robotCartesianSpaceTrajectory,
		robotJoinSpaceTrajectory,
		robotKinematics,
	)

	//start robot console control
	if err := robot_service_rest.New(&robotUseCase).Start(); err != nil {
		panic(err)
	}
}
