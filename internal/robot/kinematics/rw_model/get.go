package robot_kinematics_rw_model

import (
	"errors"
	"fmt"
	"math"
)

func (m *model) GetDegrees(point []float64) (degrees []float64, err error) {
	degrees, err = m.getDegrees(point)
	if err != nil {
		return
	}
	for i := range degrees {
		H := m.motors[i].DrumH * (degrees[i] - m.lengths[i]) / (2 * math.Pi * m.motors[i].DrumR)
		degrees[i] = (degrees[i] - H - m.lengths[i]) / m.motors[i].DrumR / math.Pi * 180
	}
	return
}

func (m *model) getDegrees(point []float64) (degrees []float64, err error) {

	// check point
	if len(point) != 3 {
		err = errors.New(fmt.Sprintf("invalid number of coordinates: %d mast be 3", len(point)))
		return
	}

	for i := range m.motors {
		cosB := (point[0] - m.motors[i].ExitPoint[0]) /
			(math.Sqrt(math.Pow(point[0]-m.motors[i].ExitPoint[0], 2) + math.Pow(point[1]-m.motors[i].ExitPoint[1], 2)))
		sinB := (point[1] - m.motors[i].ExitPoint[1]) /
			(math.Sqrt(math.Pow(point[0]-m.motors[i].ExitPoint[0], 2) + math.Pow(point[1]-m.motors[i].ExitPoint[1], 2)))

		Cs := []float64{
			m.motors[i].ExitPoint[0] + m.motors[i].RollerR*cosB,
			m.motors[i].ExitPoint[1] + m.motors[i].RollerR*sinB,
			m.motors[i].ExitPoint[2]}

		cosE := m.motors[i].RollerR /
			math.Sqrt(math.Pow(point[0]-Cs[0], 2)+math.Pow(point[1]-Cs[1], 2)+math.Pow(point[2]-Cs[2], 2))

		cosD := math.Sqrt(math.Pow(point[0]-Cs[0], 2)+math.Pow(point[1]-Cs[1], 2)) /
			math.Sqrt(math.Pow(point[0]-Cs[0], 2)+math.Pow(point[1]-Cs[1], 2)+math.Pow(point[2]-Cs[2], 2))

		gamma := math.Acos(cosE) + math.Acos(cosD)

		B := []float64{
			Cs[0] + m.motors[i].RollerR*math.Cos(gamma)*cosB,
			Cs[1] + m.motors[i].RollerR*math.Cos(gamma)*sinB,
			Cs[2] + m.motors[i].RollerR*math.Sin(gamma)}

		degrees = append(degrees,
			m.motors[i].RollerR*(math.Pi-gamma)+
				math.Sqrt(math.Pow(point[0]-B[0], 2)+math.Pow(point[1]-B[1], 2)+math.Pow(point[2]-B[2], 2)))
	}
	return
}
