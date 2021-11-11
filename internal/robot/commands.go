package robot

import (
	"errors"
	"log"
	"strconv"
	"time"
)

func (u *UseCase) execute() {
	func() {
		for f := range u.commands {
			f()
		}
	}()
}

func (u *UseCase) Stop(c chan string) {
	u.movingMutex.Lock()
	//err := uc.extruder.Stop()
	//if err != nil {
	//	c <- err.Error()
	//}
	c <- "robot pause"
	u.stopFlag = true
}

func (u *UseCase) Resume(c chan string) {
	if u.stopFlag {
		u.stopFlag = false
		u.movingMutex.Unlock()
		//err := uc.extruder.Start()
		//if err != nil {
		//	c <- err.Error()
		//}
		c <- "robot resume"
	}
}

func (u *UseCase) Reset(c chan string) {
	u.resetFlag = true
	defer func() {
		u.resetFlag = false
	}()

	if u.stopFlag {
		u.movingMutex.Unlock()
	}
	u.stopFlag = false
	go func() {
		if err := u.controller.Reset(); err == nil {
			u.position = []float64{0, 0, 0}
		} else {
			c <- err.Error()
			return
		}
		c <- "reset done"
	}()
}

func (u *UseCase) sendTrajectory(positions, degrees [][]float64, extruderSpeeds []float64, print bool) (err error) {

	count := len(degrees)
	if len(positions) != count && count != len(extruderSpeeds) {
		return errors.New("invalid arrays shape: " +
			strconv.Itoa(len(positions)) + " " +
			strconv.Itoa(len(degrees)) + " " +
			strconv.Itoa(len(extruderSpeeds)) + " ")
	}
	if count > 0 {
		if err := u.controller.SetDegrees(degrees[0]); err != nil {
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
				if counter < count && !u.resetFlag {
					go func() {
						if err := u.controller.SetDegrees(degrees[counter]); err != nil {
							log.Println("sendTrajectory: " + err.Error())
						} else {
							u.position = positions[counter]
						}
						u.movingMutex.Unlock()
					}()
					//if print {
					//	go func() {
					//		if err := uc.extruder.SetSpeed(extruderSpeeds[counter]); err != nil {
					//			log.Println("sendTrajectory: " + err.Error())
					//		}
					//	}()
					//}
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

func (u *UseCase) sendInitialiseTrajectory(degrees [][]float64, point []float64) (err error) {

	count := len(degrees)
	if count > 0 {
		if err := u.controller.SetDegrees(degrees[0]); err != nil {
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
				if counter < count && !u.resetFlag {
					go func(c int) {
						if err := u.controller.SetDegrees(degrees[c]); err != nil {
							log.Println("sendTrajectory: " + err.Error())
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
	if counter == count {
		u.position = point
	}
	return
}
