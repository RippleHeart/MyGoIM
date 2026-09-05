package main

import (
	"log"
	"mygoim/API"
	"mygoim/Chat"
	"mygoim/Conf"
	"mygoim/DB"
)

func main() {
	Conf.LoadConfig()
	DB.InitMySQL()
	DB.InitRedis()
	Chat.InitHub()
	Chat.InitMQ()
	API.InitRouters()
	go func() {
		err := API.UserEngine.Run(Conf.UserAddr)
		if err != nil {
			log.Fatal("Engine.Run: ", err)
		}
	}()
	err := API.AdminEngine.Run(Conf.AdminAddr)
	if err != nil {
		log.Fatal("Engine.Run: ", err)
	}
}
