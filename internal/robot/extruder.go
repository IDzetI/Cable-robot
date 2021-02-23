package robot

func (u *UseCase) SetExtruderSpeed(v float64) (err error) {
	return u.extruder.SetSpeed(v)
}
