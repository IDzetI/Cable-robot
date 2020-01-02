package robot_trajectory

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
func Line(pCurrent, pFinish []float64,
	speed, acceleration, deceleration, minSpeed, period float64) (points [][]float64) {

	//remaining distance
	S := utils.VectorLength(utils.VectorSub(&pCurrent, &pFinish))

	// set speeds, acceleration and deceleration
	var acc []float64
	var dec []float64
	var v []float64
	var minV []float64
	for i := 0; i < utils.MinLength(&pCurrent, &pFinish); i++ {
		k := (pFinish[i] - pCurrent[i]) / S
		acc = append(acc, acceleration*k)
		dec = append(dec, -deceleration*k)
		v = append(v, speed*k)
		minV = append(minV, minSpeed*k)
	}

	//init helped variables
	var cSpeed *[]float64
	stage := 0

	for S > e {

		if stage == 0 {
			cSpeed = utils.VectorSum(cSpeed, utils.VectorScalarMul(&period, &acc))
		}

		moduleOfCurSpeed := utils.VectorLength(cSpeed)

		if moduleOfCurSpeed > speed {
			cSpeed = &v
			stage = 1
		}

		if S <= moduleOfCurSpeed*moduleOfCurSpeed/2/deceleration {
			stage = 2
		}
		if stage == 2 {
			cSpeed = utils.VectorSum(cSpeed, utils.VectorScalarMul(&period, &dec))
		}

		if utils.VectorLength(cSpeed) < minSpeed && stage == 2 {
			cSpeed = &minV
			stage = 3
		}

		pCurrent = *utils.VectorSum(&pCurrent, utils.VectorScalarMul(&period, cSpeed))
		points = append(points, pCurrent)

		if s := utils.VectorLength(utils.VectorSub(&pCurrent, &pFinish)); s > S {
			break
		} else {
			S = s
		}
	}

	//add last point
	points = append(points, pFinish)
	return
}
