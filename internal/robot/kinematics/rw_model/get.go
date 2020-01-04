package robot_kinematics_rw_model

import (
	"errors"
	"fmt"
	"math"
)

func (m *model) GetDegrees(point []float64) (lengths []float64, err error) {
	degrees, err := m.getDegree(point)
	if err != nil {
		return
	}
	for i := range degrees {
		H := m.h * (degrees[i] - m.lengths[i]) / (2 * math.Pi * m.r)
		lengths = append(lengths, (degrees[i]-H-degrees[i])/m.R/math.Pi*180)
	}
	return
}

func (m *model) getDegree(point []float64) (degree []float64, err error) {

	// check point
	if len(point) != 3 {
		err = errors.New(fmt.Sprintf("invalid number of coordinates: %d mast be 3", len(point)))
		return
	}

	for _, C := range m.C {
		cosB := (point[0] - C[0]) /
			(math.Sqrt(math.Pow(point[0]-C[0], 2) + math.Pow(point[1]-C[1], 2)))
		sinB := (point[1] - C[1]) /
			(math.Sqrt(math.Pow(point[0]-C[0], 2) + math.Pow(point[1]-C[1], 2)))

		Cs := []float64{
			C[0] + m.r*cosB,
			C[1] + m.r*sinB,
			C[2]}

		cosE := m.r /
			math.Sqrt(math.Pow(point[0]-Cs[0], 2)+math.Pow(point[1]-Cs[1], 2)+math.Pow(point[2]-Cs[2], 2))

		cosD := math.Sqrt(math.Pow(point[0]-Cs[0], 2)+math.Pow(point[1]-Cs[1], 2)) /
			math.Sqrt(math.Pow(point[0]-Cs[0], 2)+math.Pow(point[1]-Cs[1], 2)+math.Pow(point[2]-Cs[2], 2))

		gamma := math.Acos(cosE) + math.Acos(cosD)

		B := []float64{
			Cs[0] + m.r*math.Cos(gamma)*cosB,
			Cs[1] + m.r*math.Cos(gamma)*sinB,
			Cs[2] + m.r*math.Sin(gamma)}

		degree = append(degree,
			m.r*(math.Pi-gamma)+
				math.Sqrt(math.Pow(point[0]-B[0], 2)+math.Pow(point[1]-B[1], 2)+math.Pow(point[2]-B[2], 2)))
	}
	return
}
