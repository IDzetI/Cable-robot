package robot_controller

type Controller interface {
	GetDegrees() (degrees []float64, err error)
	GetPeriod() (period float64)

	SetDegrees(degrees []float64) (err error)

	ControlON() (err error)
	ControlOFF() (err error)

	Reset() (err error)
}
