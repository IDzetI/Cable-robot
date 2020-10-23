package robot_controller_test

import (
	"github.com/IDzetI/Cable-robot/internal/robot/controller"
	"log"
)

type controller struct {
}

func New() (c robot_controller.Controller, err error) {
	log.Println("controller init")
	c = &controller{}
	return
}
