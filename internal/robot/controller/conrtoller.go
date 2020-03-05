package robot_controller

type Controller interface {
	GetDegrees() (degrees []float64, err error)

	SendTrajectory(degrees [][]float64) (err error)
	SetDegrees(degrees []float64) (err error)

	ControlON() (err error)
	ControlOFF() (err error)

	HasError() (e bool, err error)
}
