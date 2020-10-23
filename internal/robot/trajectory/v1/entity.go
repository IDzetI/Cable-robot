package robot_trajectory_v1

import (
	"github.com/IDzetI/Cable-robot/internal/robot/trajectory"
)

type trajectory struct {
	speed        float64
	minSpeed     float64
	acceleration float64
	deceleration float64

	period float64

	workspace [][]float64
}

func New(speed, minSpeed, acceleration, deceleration, period float64, workspace [][]float64) (t robot_trajectory.Trajectory, err error) {

	err = checkBoarders(workspace)
	if err != nil {
		return
	}

	t = &trajectory{
		speed:        speed,
		minSpeed:     minSpeed,
		acceleration: acceleration,
		deceleration: deceleration,
		period:       period,
		workspace:    workspace,
	}

	return
}

func (t *trajectory) SetBoarders(boarders [][]float64) (err error) {
	err = checkBoarders(boarders)
	if err != nil {
		return
	}
	t.workspace = boarders
	return
}

func (t *trajectory) GetBoarders() (boarders [][]float64) {
	return t.workspace
}

func (t *trajectory) SetSpeed(speed float64) (err error) {
	t.speed = speed
	return
}

func (t *trajectory) GetSpeed() (speed float64) {
	return t.speed
}

func (t *trajectory) SetMinSpeed(speed float64) (err error) {
	t.minSpeed = speed
	return
}

func (t *trajectory) GetMinSpeed() (speed float64) {
	return t.minSpeed
}

func (t *trajectory) SetAcceleration(acceleration float64) (err error) {
	t.acceleration = acceleration
	return
}

func (t *trajectory) GetAcceleration() (acceleration float64) {
	return t.acceleration
}

func (t *trajectory) SetDeceleration(deceleration float64) (err error) {
	t.deceleration = deceleration
	return
}

func (t *trajectory) GetDeceleration() (deceleration float64) {
	return t.deceleration
}
