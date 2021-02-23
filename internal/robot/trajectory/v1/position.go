package robot_trajectory_v1

import (
	"errors"
	"github.com/IDzetI/Cable-robot/pkg/utils"
)

func (t *trajectory) ToPoint(position, newPosition, speed []float64) (points, speeds [][]float64, err error) {
	err = t.checkPosition(position)
	if err != nil {
		return
	}
	if len(position) != len(speed) {
		err = errors.New("invalid speed dimension")
		return
	}

	//remaining distance
	S := utils.VectorLength(utils.VectorSub(&position, &newPosition))

	stage := 0
	for S > e {
		var vm, acc, dec, minV []float64
		for i := 0; i < len(position); i++ {
			k := (newPosition[i] - position[i]) / S
			acc = append(acc, t.acceleration*k)
			dec = append(dec, -t.deceleration*k)
			vm = append(vm, t.speed*k)
			minV = append(minV, t.minSpeed*k)
		}

		if stage == 0 {
			speed = *utils.VectorSum(&speed, utils.VectorScalarMul(&t.period, utils.VectorMinValue(utils.VectorSub(&speed, &vm), &acc)))
		}

		moduleOfCurSpeed := utils.VectorLength(&speed)

		if moduleOfCurSpeed > t.speed {
			speed = vm
			stage = 1
		}

		if S <= moduleOfCurSpeed*moduleOfCurSpeed/2/t.deceleration {
			stage = 2
		}
		if stage == 2 {
			speed = *utils.VectorSum(&speed, utils.VectorScalarMul(&t.period, &dec))
		}

		if utils.VectorLength(&speed) < t.minSpeed && stage == 2 {
			speed = minV
			stage = 3
		}

		position = *utils.VectorSum(&position, utils.VectorScalarMul(&t.period, &speed))

		if s := utils.VectorLength(utils.VectorSub(&position, &newPosition)); s < e || (s > S && stage > 1) {
			break
		} else {
			points = append(points, position)
			speeds = append(speeds, speed)
			S = s
		}
	}

	//add last point
	if last := len(points) - 1; points != nil && last >= 0 && speeds != nil && len(speeds) == len(points) {
		points[last] = newPosition
		speeds[last] = *utils.VectorLike(&speed, 0)
	}

	return
}
