package robot

import (
	"errors"
	"log"
	"strconv"
	"time"
)

func (uc *UseCase) execute() {
	func() {
		for f := range uc.commands {
			f()
		}
	}()
}

func (uc *UseCase) Stop(c chan string) {
	uc.movingMutex.Lock()
	err := uc.extruder.Stop()
	if err != nil {
		c <- err.Error()
	}
	uc.stopFlag = true
}

func (uc *UseCase) Resume(c chan string) {
	if uc.stopFlag {
		uc.stopFlag = false
		uc.movingMutex.Unlock()
		err := uc.extruder.Start()
		if err != nil {
			c <- err.Error()
		}
	}
}

func (uc *UseCase) Reset(c chan string) {
	uc.resetFlag = true
	if uc.stopFlag {
		uc.movingMutex.Unlock()
	}
	uc.stopFlag = false
	go func() {
		time.Sleep(time.Second)
		uc.resetFlag = false
		c <- "reset done"
	}()
}

func (uc *UseCase) sendTrajectory(positions, degrees [][]float64, extruderSpeeds []float64, print bool) (err error) {

	count := len(degrees)
	if len(positions) != count && count != len(extruderSpeeds) {
		return errors.New("invalid arrays shape: " +
			strconv.Itoa(len(positions)) + " " +
			strconv.Itoa(len(degrees)) + " " +
			strconv.Itoa(len(extruderSpeeds)) + " ")
	}
	if count > 0 {
		if err := uc.controller.SetDegrees(degrees[0]); err != nil {
			return err
		}
	}

	if err := uc.controller.ControlON(); err != nil {
		return err
	}

	//create timer
	done := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(uc.controller.GetPeriod()*1e+6) * time.Microsecond)

	// current line
	counter := 0
	go func() {
		for {
			select {
			case <-ticker.C:
				counter++
				uc.movingMutex.Lock()
				if counter < count && !uc.resetFlag {

					go func() {
						if err := uc.controller.SetDegrees(degrees[counter]); err != nil {
							log.Println("sendTrajectory: " + err.Error())
						} else {
							uc.position = positions[counter]
						}
						uc.movingMutex.Unlock()
					}()
					if print {
						go func() {
							if err := uc.extruder.SetSpeed(extruderSpeeds[counter]); err != nil {
								log.Println("sendTrajectory: " + err.Error())
							}
						}()
					}
				} else {
					uc.movingMutex.Unlock()
					done <- true
					return
				}
			}
		}
	}()
	<-done
	return
}

func (uc *UseCase) sendInitialiseTrajectory(degrees [][]float64, point []float64) (err error) {

	count := len(degrees)
	if count > 0 {
		if err := uc.controller.SetDegrees(degrees[0]); err != nil {
			return err
		}
	}

	if err := uc.controller.ControlON(); err != nil {
		return err
	}

	//create timer
	done := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(uc.controller.GetPeriod()*1e+6) * time.Microsecond)

	// current line
	counter := 0
	go func() {
		for {
			select {
			case <-ticker.C:
				counter++
				uc.movingMutex.Lock()
				if counter < count && !uc.resetFlag {
					go func() {
						if err := uc.controller.SetDegrees(degrees[counter]); err != nil {
							log.Println("sendTrajectory: " + err.Error())
						}
						uc.movingMutex.Unlock()
					}()
				} else {
					uc.movingMutex.Unlock()
					done <- true
					return
				}
			}
		}
	}()
	<-done
	if counter == count {
		uc.position = point
	}
	return
}
