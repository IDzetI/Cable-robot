package robot_service

import "github.com/IDzetI/Cable-robot/internal/robot"

type Service interface {
	Start(uc robot.UseCase) (err error)
}
