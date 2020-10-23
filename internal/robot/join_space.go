package robot

import "github.com/IDzetI/Cable-robot/pkg/utils"

func (uc *UseCase) SetSpeedJoinSpace(v float64) (err error) {
	return uc.trajectoryJoinSpace.SetSpeed(v)
}

func (uc *UseCase) SetMinSpeedJoinSpace(v float64) (err error) {
	return uc.trajectoryJoinSpace.SetMinSpeed(v)
}

func (uc *UseCase) SetAccelerationJoinSpace(v float64) (err error) {
	return uc.trajectoryJoinSpace.SetAcceleration(v)
}

func (uc *UseCase) SetDecelerationJoinSpace(v float64) (err error) {
	return uc.trajectoryJoinSpace.SetDeceleration(v)
}

func (uc *UseCase) MoveInJoinSpace(point []float64, c chan string) (err error) {

	//get current degree
	lengths, err := uc.controller.GetDegrees()
	if err != nil {
		return
	}

	//shift point
	uc.shiftPoint(&point)

	//get end degree position
	endLengths, err := uc.kinematics.GetDegrees(point)
	if err != nil {
		return
	}

	//calculate trajectory
	trajectory, _, err := uc.trajectoryJoinSpace.Line(lengths, endLengths)
	if err != nil {
		return
	}

	//execute trajectory
	c <- "robot move to " + utils.ToString(point)
	err = uc.sendInitialiseTrajectory(trajectory, point)
	c <- "robot stop"
	return
}
