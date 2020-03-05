package robot_controller_omron

import (
	"github.com/IDzetI/Cable-robot/pkg/fins"
)

type controller struct {
	robot *fins.Client
}

func Init(address string) (controller controller, err error) {
	client, err := fins.NewClient(address)
	if err != nil {
		return
	}
	controller.setClient(client)
	return
}

func (c *controller) setClient(client *fins.Client) {
	c.robot = client
}
