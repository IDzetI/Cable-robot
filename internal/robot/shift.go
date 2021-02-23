package robot

import (
	"errors"
	"strconv"
)

func (u *UseCase) SetShift(shift []float64) (err error) {
	if len(shift) != 3 {
		return errors.New("invalid shift shape: must 3 get " + strconv.Itoa(len(shift)))
	}
	u.shift = shift
	return
}

func (u *UseCase) GetShift() (shift []float64) {
	return u.shift
}

func (u *UseCase) shiftPoint(p *[]float64) {
	if u.shift != nil && p != nil && len(*p) >= len(u.shift) {
		for i := range u.shift {
			(*p)[i] += u.shift[i]
		}
	}
}
