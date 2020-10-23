package robot

import (
	"errors"
	"fmt"
	"github.com/IDzetI/Cable-robot/pkg/utils"
)

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
	if uc.position == nil {
		return errors.New("please initialize robot position")
	}
	uc.commands <- uc.moveInCartesianSpace(point, c)
	return
}

func (uc *UseCase) moveInCartesianSpace(point []float64, c chan string) func() {
	return func() {

		//check print condition
		printing := len(point) > 3 && point[3] == 1
		point = point[:3]

		//shift point
		uc.shiftPoint(&point)

		fmt.Println("move from", uc.position, "to", point)
		//get trajectory
		trajectory, extruderSpeed, err := uc.trajectoryCartesianSpace.Line(uc.position, point)
		if err != nil {
			c <- err.Error()
			return
		}

		//get degrees sequence
		var degreesTrajectory [][]float64
		for _, p := range trajectory {
			var degrees []float64
			degrees, err = uc.kinematics.GetDegrees(p)
			if err != nil {
				return
			}
			degreesTrajectory = append(degreesTrajectory, degrees)
		}

		c <- "robot move to " + utils.ToString(point)
		err = uc.sendTrajectory(trajectory, degreesTrajectory, extruderSpeed, printing)
		if err != nil {
			c <- err.Error()
			return
		}
		c <- "robot stop"
	}
}
