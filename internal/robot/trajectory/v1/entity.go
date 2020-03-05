package robot_trajectory_v1

type trajectory struct {
	speed        float64
	minSpeed     float64
	acceleration float64
	deceleration float64

	period float64

	position  []float64
	workspace [][]float64
}

func New(speed, minSpeed, acceleration, deceleration, period float64, workspace [][]float64) (t *trajectory, err error) {

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

func (t *trajectory) SetPosition(position []float64) (err error) {
	if t == nil {
		err = errorEmptyObj()
		return
	}

	err = t.checkPosition(position)
	if err != nil {
		return
	}
	t.position = position
	return
}

func (t *trajectory) GetPosition() (position []float64) {
	if t != nil {
		return t.position
	}
	return
}

func (t *trajectory) SetBoarders(boarders [][]float64) (err error) {
	if t == nil {
		err = errorEmptyObj()
		return
	}

	err = checkBoarders(boarders)
	if err != nil {
		return
	}
	t.workspace = boarders
	return
}

func (t *trajectory) GetBoarders() (boarders [][]float64) {
	if t != nil {
		return t.workspace
	}
	return
}

func (t *trajectory) SetSpeed(speed float64) (err error) {
	if t == nil {
		err = errorEmptyObj()
		return
	}

	t.speed = speed
	return
}

func (t *trajectory) GetSpeed() (speed float64) {
	if t != nil {
		return t.speed
	}
	return
}

func (t *trajectory) SetMinSpeed(speed float64) (err error) {
	if t == nil {
		err = errorEmptyObj()
		return
	}

	t.minSpeed = speed
	return
}

func (t *trajectory) GetMinSpeed() (speed float64) {
	if t != nil {
		return t.minSpeed
	}
	return
}

func (t *trajectory) SetAcceleration(acceleration float64) (err error) {
	if t == nil {
		err = errorEmptyObj()
		return
	}

	t.acceleration = acceleration
	return
}

func (t *trajectory) GetAcceleration() (acceleration float64) {
	if t != nil {
		return t.acceleration
	}
	return
}

func (t *trajectory) SetDeceleration(deceleration float64) (err error) {
	if t == nil {
		err = errorEmptyObj()
		return
	}

	t.deceleration = deceleration
	return
}

func (t *trajectory) GetDeceleration() (deceleration float64) {
	if t != nil {
		return t.deceleration
	}
	return
}
