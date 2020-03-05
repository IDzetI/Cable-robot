package main

import (
	"github.com/IDzetI/Cable-robot/config"
	"github.com/IDzetI/Cable-robot/internal/robot"
	"github.com/IDzetI/Cable-robot/internal/robot/controller/extruder/smsd"
	"github.com/IDzetI/Cable-robot/internal/robot/controller/test"
	"github.com/IDzetI/Cable-robot/internal/robot/kinematics/rw_model"
	robot_service_console "github.com/IDzetI/Cable-robot/internal/robot/service/console"
	"github.com/IDzetI/Cable-robot/internal/robot/trajectory/v1"
	"log"
)

func main() {
	log.Println("Program starting...")

	conf, err := config.Init()
	if err != nil {
		panic(err)
	}

	robotController, err := robot_controller_test.New()
	if err != nil {
		panic(err)
	}

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

	robotJoinSpaceTrajectory, err := robot_trajectory_v1.New(
		conf.JoinSpace.Speed,
		conf.JoinSpace.MinSpeed,
		conf.JoinSpace.Acceleration,
		conf.JoinSpace.Deceleration,
		conf.Period,
		conf.Workspace,
	)
	if err != nil {
		panic(err)
	}

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

	robotExtruder, err := robot_extruder_smsd.New()
	if err != nil {
		return
	}

	robotUseCase := robot.New(
		robotController,
		robotCartesianSpaceTrajectory,
		robotJoinSpaceTrajectory,
		robotKinematics,
		robotExtruder,
	)

	if err := robot_service_console.Start(robotUseCase); err != nil {
		panic(err)
	}
}
