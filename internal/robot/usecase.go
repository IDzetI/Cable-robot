package robot

type UseCase struct {
}

func (uc UseCase) ReadDegree() (degree []float64) {
	return
}

func (uc UseCase) SetSpeed(v float64) (err error) {
	return
}

func (uc UseCase) SetMinSpeed(v float64) (err error) {
	return
}

func (uc UseCase) SetAcceleration(v float64) (err error) {
	return
}

func (uc UseCase) SetDeceleration(v float64) (err error) {
	return
}

func (uc UseCase) MoveInCartesianSpace(point []float64) (err error) {
	return
}

func (uc UseCase) MoveInJoinSpace(point []float64) (err error) {
	return
}

func (uc UseCase) ControlOn() (err error) {
	return
}

func (uc UseCase) ControlOff() (err error) {
	return
}

func (uc UseCase) FileParse(file string) (err error) {
	return
}

func (uc UseCase) FileLoad(file string) (err error) {
	return
}

func (uc UseCase) FileInit() (err error) {
	return
}

func (uc UseCase) FileNext() (err error) {
	return
}

func (uc UseCase) FileCurrent() (err error) {
	return
}

func (uc UseCase) FilePrevious() (err error) {
	return
}

func (uc UseCase) FileSetCursor(i int) (err error) {
	return
}

func (uc UseCase) FileGo() (err error) {
	return
}

func (uc UseCase) FileRun() (err error) {
	return
}
