package robot_kinematics_rw_model

import robot_kinematics "github.com/IDzetI/Cable-robot/internal/robot/kinematics"

type model struct {
	lengths []float64
	motors  []Motor
}

type Motor struct {
	DrumH     float64
	DrumR     float64
	RollerR   float64
	ExitPoint []float64
}

func New(motors []Motor) (k robot_kinematics.Kinematics, err error) {
	m := model{
		motors: motors,
	}
	m.lengths, err = m.getDegrees([]float64{0, 0, 0})
	return &m, err
}
