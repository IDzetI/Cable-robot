package robot_extruder_smsd

import (
	"github.com/tarm/serial"
)

type extruder struct {
	porn *serial.Port
}

func New(p string) (e extruder, err error) {
	c := &serial.Config{Name: p, Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		return
	}
	e = extruder{porn: s}
	return
}
