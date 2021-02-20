package robot_extruder_smsd

import (
	"log"
	"time"
)

func (e extruder) SetSpeed(speed float64) (err error) {
	log.Println(time.Now(), "extrude speed", speed)
	//panic("implement me")
	return
}

func (e extruder) Stop() (err error) {
	log.Println(time.Now(), "extrude stop")
	return
}

func (e extruder) Start() (err error) {
	log.Println(time.Now(), "extrude stop")
	return
}
