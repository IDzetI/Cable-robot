package robot_trajectory_v1

import "errors"

const (
	e = 1e-3

	errorObjectTrajectoryIsEmpty = "trajectory object is empty"
)

func errorEmptyObj() error {
	return errors.New(errorObjectTrajectoryIsEmpty)
}
