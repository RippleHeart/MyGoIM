package main

import (
	"mygoim/Chat"
	"mygoim/Conf"
	"mygoim/DB"
	"mygoim/Routers"
)

func main() {
	Conf.LoadConfig()
	DB.InitMySQL()
	DB.InitRedis()
	Chat.InitHub()
	Chat.InitMQ()
	Routers.InitRouters()
}
