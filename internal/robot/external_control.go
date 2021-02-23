package robot

import (
	"errors"
	"log"
	"strconv"
	"time"
)

func (u *UseCase) ExternalControl(position []float64) (err error) {

	if u.position == nil {
		return errors.New("please initialize robot position")
	}

	//shift point
	u.shiftPoint(&position)

	t := time.Now().UnixNano()
	u.externalInterception = t

	trajectory, speed, err := u.trajectoryCartesianSpace.ToPoint(u.position, position, u.GetSpeed())
	if err != nil {
		return
	}

	//get degrees sequence
	var degreesTrajectory [][]float64
	for _, p := range trajectory {
		var degrees []float64
		degrees, err = u.kinematics.GetDegrees(p)
		if err != nil {
			return
		}
		degreesTrajectory = append(degreesTrajectory, degrees)
	}

	count := len(degreesTrajectory)
	if len(trajectory) != count || count != len(speed) || len(degreesTrajectory) == 0 {
		return errors.New("invalid arrays shape: " +
			strconv.Itoa(len(trajectory)) + " " +
			strconv.Itoa(len(speed)) + " " +
			strconv.Itoa(len(degreesTrajectory)) + " ")
	}
	if count > 0 {
		if err := u.controller.SetDegrees(degreesTrajectory[0]); err != nil {
			return err
		}
	}

	if err := u.controller.ControlON(); err != nil {
		return err
	}

	//create timer
	done := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(u.controller.GetPeriod()*1e+6) * time.Microsecond)

	// current line
	counter := 0
	go func() {
		for {
			select {
			case <-ticker.C:
				counter++
				u.movingMutex.Lock()
				if counter < count && !u.resetFlag && u.externalInterception <= t {
					go func(counter int) {
						if err := u.controller.SetDegrees(degreesTrajectory[counter]); err != nil {
							log.Println("sendTrajectory: " + err.Error())
						} else {
							u.position = trajectory[counter]
						}
						u.movingMutex.Unlock()
					}(counter)
				} else {
					u.movingMutex.Unlock()
					done <- true
					return
				}
			}
		}
	}()
	<-done
	return
}
