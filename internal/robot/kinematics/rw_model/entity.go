package robot_kinematics_rw_model

type model struct {
	h       float64
	r       float64
	R       float64
	lengths []float64
	C       [][]float64
}

func Init(h float64, r float64, R float64, C [][]float64) (m model, err error) {
	m = model{
		h: h,
		r: r,
		R: R,
		C: C,
	}
	m.lengths, err = m.getDegrees([]float64{0, 0, 0})
	return
}
