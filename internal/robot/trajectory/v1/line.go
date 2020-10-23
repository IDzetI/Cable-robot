package robot_trajectory_v1

import (
	"github.com/IDzetI/Cable-robot/pkg/utils"
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
func (t *trajectory) Line(position, newPosition []float64) (points [][]float64, extruderSpeed []float64, err error) {
	err = t.checkPosition(newPosition)
	if err != nil {
		return
	}

	//remaining distance
	S := utils.VectorLength(utils.VectorSub(&position, &newPosition))

	// set speeds, acceleration and deceleration
	var acc []float64
	var dec []float64
	var v []float64
	var minV []float64
	var cSpeed []float64
	for i := 0; i < utils.MinLength(&position, &newPosition); i++ {
		k := (newPosition[i] - position[i]) / S
		acc = append(acc, t.acceleration*k)
		dec = append(dec, -t.deceleration*k)
		v = append(v, t.speed*k)
		minV = append(minV, t.minSpeed*k)
		cSpeed = append(cSpeed, 0)
	}

	//init helped variables
	stage := 0
	for S > e {

		if stage == 0 {

			cSpeed = *utils.VectorSum(&cSpeed, utils.VectorScalarMul(&t.period, &acc))
		}

		moduleOfCurSpeed := utils.VectorLength(&cSpeed)

		if moduleOfCurSpeed > t.speed {
			cSpeed = v
			stage = 1
		}

		if S <= moduleOfCurSpeed*moduleOfCurSpeed/2/t.deceleration {
			stage = 2
		}
		if stage == 2 {
			cSpeed = *utils.VectorSum(&cSpeed, utils.VectorScalarMul(&t.period, &dec))
		}

		if utils.VectorLength(&cSpeed) < t.minSpeed && stage == 2 {
			cSpeed = minV
			stage = 3
		}

		position = *utils.VectorSum(&position, utils.VectorScalarMul(&t.period, &cSpeed))

		if s := utils.VectorLength(utils.VectorSub(&position, &newPosition)); s > S {
			break
		} else {
			points = append(points, position)
			extruderSpeed = append(extruderSpeed, utils.VectorLength(&cSpeed)/t.speed)
			S = s
		}

	}

	//add last point
	if last := len(points) - 1; points != nil && last >= 0 && extruderSpeed != nil && len(extruderSpeed) == len(points) {
		points[last] = newPosition
		extruderSpeed[last] = 0
	}

	return
}
