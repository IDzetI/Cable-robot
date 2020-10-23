package robot_extruder

type Extruder interface {
	SetSpeed(speed float64) (err error)
	Stop() (err error)
	Start() (err error)
}
