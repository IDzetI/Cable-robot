package robot

func (uc *UseCase) SetExtruderSpeed(v float64) (err error) {
	return uc.extruder.SetSpeed(v)
}
