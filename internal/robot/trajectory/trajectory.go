package robot_trajectory

type Trajectory interface {
	//Init(speed, minSpeed, acceleration, deceleration, period float64, position[]float64, boarders[][]float64)(err error)
	Line(position, newPosition []float64) (points [][]float64, extruderSpeed []float64, err error)

	SetBoarders(boarders [][]float64) (err error)
	GetBoarders() (boarders [][]float64)

	SetSpeed(speed float64) (err error)
	GetSpeed() (speed float64)

	SetMinSpeed(speed float64) (err error)
	GetMinSpeed() (speed float64)

	SetAcceleration(acceleration float64) (err error)
	GetAcceleration() (acceleration float64)

	SetDeceleration(deceleration float64) (err error)
	GetDeceleration() (deceleration float64)
}
