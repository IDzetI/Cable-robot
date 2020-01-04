package robot

func (uc *UseCase) SetSpeedCartesianSpace(v float64) (err error) {
	return uc.trajectoryCartesianSpace.SetSpeed(v)
}

func (uc *UseCase) SetMinSpeedCartesianSpace(v float64) (err error) {
	return uc.trajectoryCartesianSpace.SetMinSpeed(v)
}

func (uc *UseCase) SetAccelerationCartesianSpace(v float64) (err error) {
	return uc.trajectoryCartesianSpace.SetAcceleration(v)
}

func (uc *UseCase) SetDecelerationCartesianSpace(v float64) (err error) {
	return uc.trajectoryCartesianSpace.SetDeceleration(v)
}

func (uc *UseCase) MoveInCartesianSpace(point []float64, c chan string) (err error) {

	//get trajectory
	points, err := uc.trajectoryCartesianSpace.Line(point)

	var trajectoryLengths [][]float64
	for _, p := range points {
		var lengths []float64
		lengths, err = uc.kinematics.GetDegrees(p)
		if err != nil {
			return
		}
		trajectoryLengths = append(trajectoryLengths, lengths)
	}

	err = uc.controller.SendTrajectory(trajectoryLengths)
	return
}
