package robot_kinematics

type Kinematics interface {
	GetDegrees(point []float64) (length []float64, err error)
}
