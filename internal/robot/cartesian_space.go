package robot

import (
	"errors"
	"fmt"
	"github.com/IDzetI/Cable-robot/pkg/utils"
)

func (u *UseCase) SetSpeedCartesianSpace(v float64) (err error) {
	return u.trajectoryCartesianSpace.SetSpeed(v)
}

func (u *UseCase) SetMinSpeedCartesianSpace(v float64) (err error) {
	return u.trajectoryCartesianSpace.SetMinSpeed(v)
}

func (u *UseCase) SetAccelerationCartesianSpace(v float64) (err error) {
	return u.trajectoryCartesianSpace.SetAcceleration(v)
}

func (u *UseCase) SetDecelerationCartesianSpace(v float64) (err error) {
	return u.trajectoryCartesianSpace.SetDeceleration(v)
}

func (u *UseCase) MoveInCartesianSpace(point []float64, c chan string) (err error) {
	if u.position == nil {
		return errors.New("please initialize robot position")
	}
	u.commands <- u.moveInCartesianSpace(point, c)
	return
}

func (u *UseCase) moveInCartesianSpace(point []float64, c chan string) func() {
	return func() {

		//check print condition
		printing := len(point) > 3 && point[3] == 1
		point = point[:3]

		//shift point
		u.shiftPoint(&point)

		fmt.Println("move from", u.position, "to", point)
		//get trajectory
		trajectory, extruderSpeed, err := u.trajectoryCartesianSpace.Line(u.position, point)
		if err != nil {
			c <- err.Error()
			return
		}

		//get degrees sequence
		var degreesTrajectory [][]float64
		for _, p := range trajectory {
			var degrees []float64
			degrees, err = u.kinematics.GetDegrees(p)
			if err != nil {
				return
			}
			degreesTrajectory = append(degreesTrajectory, degrees)
		}

		c <- "robot move to " + utils.ToString(point)
		err = u.sendTrajectory(trajectory, degreesTrajectory, extruderSpeed, printing)
		if err != nil {
			c <- err.Error()
			return
		}
		c <- "robot stop"
	}
}
