package robot

import (
	"errors"
	"strconv"
)

func (uc *UseCase) SetShift(shift []float64) (err error) {
	if len(shift) != 3 {
		return errors.New("invalid shift shape: must 3 get " + strconv.Itoa(len(shift)))
	}
	uc.shift = shift
	return
}

func (uc *UseCase) GetShift() (shift []float64) {
	return uc.shift
}

func (uc *UseCase) shiftPoint(p *[]float64) {
	if uc.shift != nil && p != nil && len(*p) >= len(uc.shift) {
		for i := range uc.shift {
			(*p)[i] += uc.shift[i]
		}
	}
}
