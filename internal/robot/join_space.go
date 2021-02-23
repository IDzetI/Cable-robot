package robot

import "github.com/IDzetI/Cable-robot/pkg/utils"

func (u *UseCase) SetSpeedJoinSpace(v float64) (err error) {
	return u.trajectoryJoinSpace.SetSpeed(v)
}

func (u *UseCase) SetMinSpeedJoinSpace(v float64) (err error) {
	return u.trajectoryJoinSpace.SetMinSpeed(v)
}

func (u *UseCase) SetAccelerationJoinSpace(v float64) (err error) {
	return u.trajectoryJoinSpace.SetAcceleration(v)
}

func (u *UseCase) SetDecelerationJoinSpace(v float64) (err error) {
	return u.trajectoryJoinSpace.SetDeceleration(v)
}

func (u *UseCase) MoveInJoinSpace(point []float64, c chan string) (err error) {

	//get current degree
	degrees, err := u.controller.GetDegrees()
	if err != nil {
		return
	}

	//shift point
	u.shiftPoint(&point)

	//get end degree position
	endDegrees, err := u.kinematics.GetDegrees(point)
	if err != nil {
		return
	}

	//calculate trajectory
	trajectory, _, err := u.trajectoryJoinSpace.Line(degrees, endDegrees)
	if err != nil {
		return
	}

	//execute trajectory
	c <- "robot move to " + utils.ToString(point)
	err = u.sendInitialiseTrajectory(trajectory, point)
	c <- "robot stop"
	return
}
