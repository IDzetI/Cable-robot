package robot

func (uc *UseCase) SetSpeedJoinSpace(v float64) (err error) {
	return uc.trajectoryJoinSpace.SetSpeed(v)
}

func (uc *UseCase) SetMinSpeed(v float64) (err error) {
	return uc.trajectoryJoinSpace.SetMinSpeed(v)
}

func (uc *UseCase) SetAcceleration(v float64) (err error) {
	return uc.trajectoryJoinSpace.SetAcceleration(v)
}

func (uc *UseCase) SetDeceleration(v float64) (err error) {
	return uc.trajectoryJoinSpace.SetDeceleration(v)
}

func (uc *UseCase) MoveInJoinSpace(point []float64, c chan string) (err error) {

	//get current degree
	lengths, err := uc.controller.GetDegrees()
	if err != nil {
		return
	}

	//set current degree position
	err = uc.trajectoryJoinSpace.SetPosition(lengths)
	if err != nil {
		return
	}

	//get end degree position
	endLengths, err := uc.kinematics.GetDegrees(point)
	if err != nil {
		return
	}

	//calculate trajectory
	trajectoryLengths, err := uc.trajectoryJoinSpace.Line(endLengths)
	if err != nil {
		return
	}

	//execute trajectory
	return uc.controller.SendTrajectory(trajectoryLengths)
}
