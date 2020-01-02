package robot_controller

type Controller interface {
	GetLengths() (lengths []float64, err error)
	SetLengths(lengths []float64) (err error)
	ControlON() (err error)
	ControlOFF() (err error)
	HasError() (e bool, err error)
}
