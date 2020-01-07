package main

import (
	"bufio"
	"fmt"
	"github.com/IDzetI/Cable-robot.git/internal/robot/controller/omron"
	"github.com/IDzetI/Cable-robot.git/pkg/configs"
	"log"
	"os"
)

func main() {
	log.Println("Program starting...")

	//Read config
	log.Println("Read robot config...")
	conf, err := configs.LoadConfig(defaultConfigPath)
	for err != nil {
		log.Println("Read config error: " + err.Error())
		fmt.Println("Please enter config file path:")
		reader := bufio.NewReader(os.Stdin)
		//read line
		path, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		conf, err = configs.LoadConfig(path)
	}

	//initialize controller
	address, err := conf.GetString(configRobotIp)
	if err != nil {
		log.Fatal(err)
	}
	_, err = robot_controller_omron.Init(address)
	if err != nil {
		log.Fatal(err)
	}
}
