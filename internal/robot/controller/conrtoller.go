package robot_controller

type Controller interface {
	GetDegrees() (lengths []float64, err error)

	SendTrajectory(lengths [][]float64) (err error)
	SetLengths(lengths []float64) (err error)

	ControlON() (err error)
	ControlOFF() (err error)

	HasError() (e bool, err error)
}
