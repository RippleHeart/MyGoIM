package main

import (
	"mygoim/Conf"
	"mygoim/DB"
	"mygoim/Routers"
	"mygoim/WS"
)

func main() {
	Conf.LoadConfig()
	DB.InitMySQL()
	DB.InitRedis()
	WS.InitHub()
	WS.InitMQ()
	Routers.InitRouters()
}
