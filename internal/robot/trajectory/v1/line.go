package robot_trajectory_v1

import (
	"github.com/IDzetI/Cable-robot.git/pkg/utils"
)

/*
current position [x y z] in mm
new position [x y z] in mm
speed in mm/s
acceleration in mm/s^2
deceleration in mm/s^2
min speed in mm/s
period in s

returns trajectory
*/
func (t *trajectory) Line(position []float64) (points [][]float64, err error) {
	if t == nil {
		err = errorEmptyObj()
		return
	}

	err = checkPosition(position, t.boarders)
	if err != nil {
		return
	}

	//remaining distance
	S := utils.VectorLength(utils.VectorSub(&t.position, &position))

	// set speeds, acceleration and deceleration
	var acc []float64
	var dec []float64
	var v []float64
	var minV []float64
	for i := 0; i < utils.MinLength(&t.position, &position); i++ {
		k := (position[i] - t.position[i]) / S
		acc = append(acc, t.acceleration*k)
		dec = append(dec, -t.deceleration*k)
		v = append(v, t.speed*k)
		minV = append(minV, t.minSpeed*k)
	}

	//init helped variables
	var cSpeed *[]float64
	stage := 0

	for S > e {

		if stage == 0 {
			cSpeed = utils.VectorSum(cSpeed, utils.VectorScalarMul(&t.period, &acc))
		}

		moduleOfCurSpeed := utils.VectorLength(cSpeed)

		if moduleOfCurSpeed > t.speed {
			cSpeed = &v
			stage = 1
		}

		if S <= moduleOfCurSpeed*moduleOfCurSpeed/2/t.deceleration {
			stage = 2
		}
		if stage == 2 {
			cSpeed = utils.VectorSum(cSpeed, utils.VectorScalarMul(&t.period, &dec))
		}

		if utils.VectorLength(cSpeed) < t.minSpeed && stage == 2 {
			cSpeed = &minV
			stage = 3
		}

		t.position = *utils.VectorSum(&t.position, utils.VectorScalarMul(&t.period, cSpeed))
		points = append(points, t.position)

		if s := utils.VectorLength(utils.VectorSub(&t.position, &position)); s > S {
			break
		} else {
			S = s
		}
	}

	//add last point
	points = append(points, position)
	return
}
