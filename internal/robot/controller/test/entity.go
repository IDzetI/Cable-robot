package robot_controller_test

import (
	robot_controller "github.com/IDzetI/Cable-robot/internal/robot/controller"
	"log"
)

type controller struct {
}

func New() (controller robot_controller.Controller, err error) {
	log.Println("controller init")
	return
}
