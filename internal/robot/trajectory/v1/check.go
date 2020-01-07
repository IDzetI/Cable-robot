package robot_trajectory_v1

import (
	"errors"
	"fmt"
	"github.com/IDzetI/Cable-robot.git/pkg/utils"
)

func checkPosition(position []float64, boarders [][]float64) (err error) {
	for i := range position {
		if position[i] < boarders[i][0] || position[i] > boarders[i][1] {
			return errors.New(
				fmt.Sprintf(
					"coordinate out of boarders: %f (%f,%f)",
					position[i],
					boarders[i][0],
					boarders[i][1],
				),
			)
		}
	}
	return
}

func checkBoarders(boarders [][]float64) (err error) {
	for _, boarder := range boarders {
		if len(boarder) != 2 || boarder[0] > boarder[1] {
			return errors.New("invalid boarders: " + utils.ToString(boarder))
		}
	}
	return
}
