package robot_trajectory_v1

import (
	"errors"
	"fmt"
	"github.com/IDzetI/Cable-robot/pkg/utils"
)

func (t *trajectory) checkPosition(position []float64) (err error) {
	if len(t.position) != 3 {
		return errors.New("please initialise robot position")
	}
	for i := range position {
		if position[i] < t.workspace[i][0] || position[i] > t.workspace[i][1] {
			return errors.New(
				fmt.Sprintf(
					"coordinate out of workspace: %f (%f,%f)",
					position[i],
					t.workspace[i][0],
					t.workspace[i][1],
				),
			)
		}
	}
	return
}

func checkBoarders(boarders [][]float64) (err error) {
	for _, boarder := range boarders {
		if len(boarder) != 2 || boarder[0] > boarder[1] {
			return errors.New("invalid workspace: " + utils.ToString(boarder))
		}
	}
	return
}
