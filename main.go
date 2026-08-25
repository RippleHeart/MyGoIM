package main

import (
	"mygoim/Conf"
	"mygoim/DBInit"
)

func main() {
	Conf.LoadConfig()
	DBInit.InitMySQl()
	DBInit.InitRedis()
}
